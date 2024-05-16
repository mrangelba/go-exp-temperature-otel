package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/mrangelba/go-exp-temperature/service-b/internal/domain/entity"
	"github.com/mrangelba/go-exp-temperature/service-b/internal/usecase"
)

type mockCEPGateway struct {
	GetFunc func(ctx context.Context, cep string) (*entity.CEP, error)
}

func (m *mockCEPGateway) Get(ctx context.Context, cep string) (*entity.CEP, error) {
	return m.GetFunc(ctx, cep)
}

type mockWeatherGateway struct {
	GetFunc func(ctx context.Context, cityName string) (*entity.Weather, error)
}

func (m *mockWeatherGateway) Get(ctx context.Context, cityName string) (*entity.Weather, error) {
	return m.GetFunc(ctx, cityName)
}

func TestWeatherUseCase_Get(t *testing.T) {
	ctx := context.Background()
	cep := "12345678"
	expectedWeather := &entity.Weather{TempC: 25.5}

	mockCEPGateway := &mockCEPGateway{
		GetFunc: func(ctx context.Context, cep string) (*entity.CEP, error) {
			if cep != "12345678" {
				t.Errorf("unexpected CEP: %s", cep)
			}
			return &entity.CEP{CityName: "your_city_name"}, nil
		},
	}

	mockWeatherGateway := &mockWeatherGateway{
		GetFunc: func(ctx context.Context, cityName string) (*entity.Weather, error) {
			if cityName != "your_city_name" {
				t.Errorf("unexpected city name: %s", cityName)
			}
			return expectedWeather, nil
		},
	}

	useCase := usecase.NewWeatherUseCase(mockCEPGateway, mockWeatherGateway)

	weather, err := useCase.Get(ctx, cep)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if weather != expectedWeather {
		t.Errorf("unexpected weather: %+v", weather)
	}
}

func TestWeatherUseCase_Get_CEPError(t *testing.T) {
	ctx := context.Background()
	cep := "12345678"
	expectedError := errors.New("CEP error")

	mockCEPGateway := &mockCEPGateway{
		GetFunc: func(ctx context.Context, cep string) (*entity.CEP, error) {
			if cep != "12345678" {
				t.Errorf("unexpected CEP: %s", cep)
			}
			return nil, expectedError
		},
	}

	mockWeatherGateway := &mockWeatherGateway{
		GetFunc: func(ctx context.Context, cityName string) (*entity.Weather, error) {
			t.Errorf("unexpected call to Get method of WeatherGateway")
			return nil, nil
		},
	}

	useCase := usecase.NewWeatherUseCase(mockCEPGateway, mockWeatherGateway)

	_, err := useCase.Get(ctx, cep)
	if err != expectedError {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestWeatherUseCase_Get_WeatherError(t *testing.T) {
	ctx := context.Background()
	cep := "12345678"
	expectedError := errors.New("weather error")

	mockCEPGateway := &mockCEPGateway{
		GetFunc: func(ctx context.Context, cep string) (*entity.CEP, error) {
			if cep != "12345678" {
				t.Errorf("unexpected CEP: %s", cep)
			}
			return &entity.CEP{CityName: "your_city_name"}, nil
		},
	}

	mockWeatherGateway := &mockWeatherGateway{
		GetFunc: func(ctx context.Context, cityName string) (*entity.Weather, error) {
			if cityName != "your_city_name" {
				t.Errorf("unexpected city name: %s", cityName)
			}
			return nil, expectedError
		},
	}

	useCase := usecase.NewWeatherUseCase(mockCEPGateway, mockWeatherGateway)

	_, err := useCase.Get(ctx, cep)
	if err != expectedError {
		t.Errorf("unexpected error: %v", err)
	}
}
