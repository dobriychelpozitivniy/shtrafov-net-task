package models

type Config struct {
	App               AppConfig
	RusporfileService RusporfileServiceConfig
}

type AppConfig struct {
	GrpcServerPort     string
	GrpcHttpServerPort string
}

type RusporfileServiceConfig struct {
	Url string
}
