package main

import (
	"log"

	"github.com/mrangelba/go-exp-temperature/service-b/internal/infra/http/rest"
	"github.com/mrangelba/go-exp-temperature/service-b/pkg/otel"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	url := "http://zipkin:9411/api/v2/spans"

	_, err := otel.NewTracer(url, "service-b")
	if err != nil {
		log.Fatalf("unable to initialize tracer provider due: %v", err)
	}

	rest.Initialize()
}
