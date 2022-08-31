package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"reflect"
	"shtrafov-net-task/internal/models"
	"strings"
)

func New(path string) (*models.Config, error) {
	var config models.Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, fmt.Errorf("error decode config file: %s", err)
	}

	indent := strings.Repeat(" ", 14)

	fmt.Print("Decoded")
	typ, val := reflect.TypeOf(config), reflect.ValueOf(config)
	for i := 0; i < typ.NumField(); i++ {
		indent := indent
		if i == 0 {
			indent = strings.Repeat(" ", 7)
		}
		fmt.Printf("%s%-11s → %v\n", indent, typ.Field(i).Name, val.Field(i).Interface())
	}

	return &config, nil
}
