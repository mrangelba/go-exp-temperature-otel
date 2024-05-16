package weatherapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mrangelba/go-exp-temperature/service-a/internal/domain"
	"github.com/mrangelba/go-exp-temperature/service-a/internal/domain/entity"
	"github.com/mrangelba/go-exp-temperature/service-a/pkg/error_handle"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

var (
	BASE_URL = "http://service-b:8081/cep"
)

type gateway struct {
	client *http.Client
}

type weather struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func NewWeatherGateway(client *http.Client) domain.WeatherGateway {
	return &gateway{
		client: client,
	}
}

func (g gateway) Get(ctx context.Context, cep string) (*entity.Weather, error) {
	tr := otel.Tracer("microservice-otel")
	ctx, span := tr.Start(ctx, "get weather from service-b")
	defer span.End()

	var weatherOutput weather

	url := fmt.Sprintf("%s/%s", BASE_URL, cep)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(nil))

	if err != nil {
		return nil, err
	}

	defer request.Body.Close()

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(request.Header))
	response, err := g.client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	println(response.StatusCode)

	if response.StatusCode == 404 {
		return nil, error_handle.ErrNotFound
	}

	if response.StatusCode == 422 {
		return nil, error_handle.ErrUnprocessableEntity
	}

	err = json.NewDecoder(response.Body).Decode(&weatherOutput)

	if err != nil {
		return nil, err
	}

	return &entity.Weather{
		TempC: weatherOutput.TempC,
		TempF: weatherOutput.TempF,
		TempK: weatherOutput.TempK,
	}, nil
}
