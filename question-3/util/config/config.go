package config

import (
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	CommonConfig CommonConfig
}

type CommonConfig struct {
	BeefApi  string `envconfig:"BEEF_API" default:"localhost"`
	HttpPort string `envconfig:"HTTP_PORT" default:"8080"`
	GrpcPort string `envconfig:"GRPC_PORT" default:"50051"`
}

func GetAppConfig() AppConfig {
	var app AppConfig
	envconfig.MustProcess("APP", &app.CommonConfig)
	return app
}
