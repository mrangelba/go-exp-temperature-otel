package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/mrangelba/go-exp-temperature/service-a/internal/domain"
	"github.com/mrangelba/go-exp-temperature/service-a/internal/dto"
	"github.com/mrangelba/go-exp-temperature/service-a/pkg/error_handle"

	"github.com/go-chi/chi/v5"
)

type controller struct {
	usecase domain.WeatherUseCase
}

func NewWeatherHandlers(usecase domain.WeatherUseCase) domain.WeatherHandlers {
	return &controller{
		usecase: usecase,
	}
}

func (controller controller) Get(response http.ResponseWriter, request *http.Request) {
	cepP := chi.URLParam(request, "cep")
	cepRegex := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	response.Header().Set("Content-Type", "application/json")

	if !cepRegex.MatchString(cepP) {
		statusCode, message := error_handle.Handle(error_handle.ErrUnprocessableEntity)
		response.WriteHeader(statusCode)
		json.NewEncoder(response).Encode(message)

		return
	}

	weather, err := controller.usecase.Get(request.Context(), cepP)

	if err != nil {
		statusCode, message := error_handle.Handle(err)
		response.WriteHeader(statusCode)
		json.NewEncoder(response).Encode(message)

		return
	}

	weatherResponse := dto.WeatherResponse{
		TempC: dto.Number(weather.TempC),
		TempF: dto.Number(weather.TempF),
		TempK: dto.Number(weather.TempK),
	}

	json.NewEncoder(response).Encode(weatherResponse)
}
