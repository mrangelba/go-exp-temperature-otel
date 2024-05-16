package domain

import (
	"context"
	"net/http"

	"github.com/mrangelba/go-exp-temperature/service-b/internal/domain/entity"
)

type WeatherGateway interface {
	Get(context.Context, string) (*entity.Weather, error)
}

type WeatherUseCase interface {
	Get(context.Context, string) (*entity.Weather, error)
}

type WeatherHandlers interface {
	Get(http.ResponseWriter, *http.Request)
}

type CEPGateway interface {
	Get(context.Context, string) (*entity.CEP, error)
}
