package main

import (
	"log"

	"github.com/mrangelba/go-exp-temperature/service-a/internal/infra/http/rest"
	"github.com/mrangelba/go-exp-temperature/service-a/pkg/otel"
)

func main() {
	url := "http://zipkin:9411/api/v2/spans"

	_, err := otel.NewTracer(url, "service-a")
	if err != nil {
		log.Fatal(err)
	}

	rest.Initialize()
}
