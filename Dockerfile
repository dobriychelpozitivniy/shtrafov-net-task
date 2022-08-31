FROM golang:alpine as builder


WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o service ./cmd/main.go


# generate clean, final image for end users
FROM alpine

COPY --from=builder /build/service /app/
COPY --from=builder /build/configs /app/configs/

EXPOSE 8083
EXPOSE 8082

# executable
ENTRYPOINT [ "/app/service", "--config", "/app/configs/config.toml" ]