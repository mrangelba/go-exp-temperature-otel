package usecase

import (
	"context"

	"github.com/mrangelba/go-exp-temperature/service-b/internal/domain"
	"github.com/mrangelba/go-exp-temperature/service-b/internal/domain/entity"
)

type usecase struct {
	cepHTTPClient     domain.CEPGateway
	weatherHTTPClient domain.WeatherGateway
}

func NewWeatherUseCase(
	cepHTTPClient domain.CEPGateway,
	weatherHTTPClient domain.WeatherGateway,
) domain.WeatherUseCase {
	return &usecase{
		cepHTTPClient:     cepHTTPClient,
		weatherHTTPClient: weatherHTTPClient,
	}
}

func (usecase usecase) Get(ctx context.Context, cep string) (*entity.Weather, error) {
	cepResponse, err := usecase.cepHTTPClient.Get(ctx, cep)

	if err != nil {
		return nil, err
	}

	weatherResponse, err := usecase.weatherHTTPClient.Get(ctx, cepResponse.CityName)

	if err != nil {
		return nil, err
	}

	return weatherResponse, nil
}
