package di

import (
	"net/http"

	"github.com/mrangelba/go-exp-temperature/service-a/internal/domain"
	"github.com/mrangelba/go-exp-temperature/service-a/internal/infra/gateway/weatherapi"
	"github.com/mrangelba/go-exp-temperature/service-a/internal/infra/http/handlers"
	"github.com/mrangelba/go-exp-temperature/service-a/internal/usecase"
)

func ConfigWebController() domain.WeatherHandlers {
	httpClient := http.DefaultClient

	weatherHttpClient := weatherapi.NewWeatherGateway(httpClient)
	weatherUseCase := usecase.NewWeatherUseCase(weatherHttpClient)

	return handlers.NewWeatherHandlers(weatherUseCase)
}
