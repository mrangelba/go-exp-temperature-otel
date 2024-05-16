package weatherapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mrangelba/go-exp-temperature/service-b/internal/domain"
	"github.com/mrangelba/go-exp-temperature/service-b/internal/domain/entity"
	"go.opentelemetry.io/otel"
)

var (
	BASE_URL = "https://api.weatherapi.com/v1/current.json"
)

type gateway struct {
	client        *http.Client
	weatherApiKey string
}

type weather struct {
	Current current `json:"current"`
}

type current struct {
	TempC float32 `json:"temp_c"`
}

func NewWeatherGateway(client *http.Client, weatherApiKey string) domain.WeatherGateway {
	return &gateway{
		client:        client,
		weatherApiKey: weatherApiKey,
	}
}

func (g gateway) Get(ctx context.Context, cityName string) (*entity.Weather, error) {
	tr := otel.Tracer("microservice-otel")
	ctx, span := tr.Start(ctx, "get weather - WeatherAPI")
	defer span.End()

	var weatherOutput weather

	params := url.Values{}
	params.Add("key", g.weatherApiKey)
	params.Add("q", cityName)
	params.Add("aqi", "no")

	url := fmt.Sprintf("%s?%s", BASE_URL, params.Encode())

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(nil))

	if err != nil {
		return nil, err
	}

	defer request.Body.Close()

	response, err := g.client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&weatherOutput)

	if err != nil {
		return nil, err
	}

	return &entity.Weather{
		TempC: weatherOutput.Current.TempC,
	}, nil
}
