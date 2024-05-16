package usecase

import (
	"context"

	"github.com/mrangelba/go-exp-temperature/service-a/internal/domain"
	"github.com/mrangelba/go-exp-temperature/service-a/internal/domain/entity"
)

type usecase struct {
	weatherHTTPClient domain.WeatherGateway
}

func NewWeatherUseCase(
	weatherHTTPClient domain.WeatherGateway,
) domain.WeatherUseCase {
	return &usecase{
		weatherHTTPClient: weatherHTTPClient,
	}
}

func (usecase usecase) Get(ctx context.Context, cep string) (*entity.Weather, error) {
	weatherResponse, err := usecase.weatherHTTPClient.Get(ctx, cep)

	if err != nil {
		return nil, err
	}

	return weatherResponse, nil
}
