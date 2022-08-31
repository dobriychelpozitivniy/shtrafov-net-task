package rusprofileService

import (
	"bytes"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"io"
	"net/http"
	grpcErr "shtrafov-net-task/internal/infrastructure/grpc-error"
	grpcService2 "shtrafov-net-task/internal/infrastructure/proto/shtrafov-net-task"
	"strings"
)

type service struct {
	grpcService2.UnimplementedINNServiceServer
	url string
}

func New(url string) *service {
	return &service{url: url}
}

func (service service) GetOrganisationInfo(ctx context.Context, req *grpcService2.INNRequest) (*grpcService2.INNResponse, error) {
	htmlBytes, err := service.getOrganisationHtml(ctx, req.INN)
	if err != nil {
		log.Error().Msgf("error while get organisation info html: %s", err)
		return nil, err
	}

	organisationInfo, err := parseOrganisationInfo(htmlBytes)
	if err != nil {
		log.Error().Msgf("error while parse organisation info: %s", err)
		return nil, err
	}

	return organisationInfo, nil
}

func (service service) getOrganisationHtml(ctx context.Context, inn string) ([]byte, error) {
	url := fmt.Sprintf(service.url, inn)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, grpcErr.New(codes.Internal, "error create request", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, grpcErr.New(codes.Internal, "error do request to external service", err)
	}

	if resp.StatusCode >= 300 {
		return nil, grpcErr.New(codes.Internal, "external service response incorrectly", nil)
	}

	isHaveIdInUrl := strings.Contains(resp.Request.URL.String(), "id")
	if !isHaveIdInUrl {
		return nil, grpcErr.New(codes.InvalidArgument, "INN not found", nil)
	}

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, grpcErr.New(codes.Internal, "error read body from response", err)
	}

	return b, nil
}

func parseOrganisationInfo(html []byte) (*grpcService2.INNResponse, error) {
	var organisationInfo = &grpcService2.INNResponse{}

	reader := bytes.NewReader(html)
	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, grpcErr.New(codes.Internal, "error create reader from html bytes", err)
	}

	organisationInfo.FIO = document.Find(".company-info__text").
		Find("a").
		Find("span").
		Text()

	organisationInfo.NAME = document.Find(".company-name").Text()

	document.Find(".company-requisites").Find(".company-row").Find(".company-col").Each(func(i int, s *goquery.Selection) {
		if i == 1 {
			s.Find(".copy_target").Each(func(i int, s *goquery.Selection) {
				switch {
				case i == 0:
					organisationInfo.INN = s.Text()
				case i == 1:
					organisationInfo.KPP = s.Text()
				}
			})
		}
	})

	return organisationInfo, nil
}
