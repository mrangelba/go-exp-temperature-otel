package di

import (
	"net/http"

	"github.com/mrangelba/go-exp-temperature/service-b/internal/domain"
	"github.com/mrangelba/go-exp-temperature/service-b/internal/infra/gateway/viacep"
	"github.com/mrangelba/go-exp-temperature/service-b/internal/infra/gateway/weatherapi"
	"github.com/mrangelba/go-exp-temperature/service-b/internal/infra/http/handlers"
	"github.com/mrangelba/go-exp-temperature/service-b/internal/usecase"
	"github.com/spf13/viper"
)

func ConfigWebController() domain.WeatherHandlers {
	httpClient := http.DefaultClient

	cepHttpClient := viacep.NewCepGateway(httpClient)
	weatherHttpClient := weatherapi.NewWeatherGateway(httpClient, viper.GetString("weather_api_key"))
	weatherUseCase := usecase.NewWeatherUseCase(cepHttpClient, weatherHttpClient)

	return handlers.NewWeatherHandlers(weatherUseCase)
}
