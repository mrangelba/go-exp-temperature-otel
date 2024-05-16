package entity_test

import (
	"testing"

	"github.com/mrangelba/go-exp-temperature/service-b/internal/domain/entity"
)

func TestWeather_TempF(t *testing.T) {
	weather := &entity.Weather{TempC: 25.5}
	expectedTempF := float32(77.9)

	tempF := weather.TempF()

	if tempF != expectedTempF {
		t.Errorf("unexpected temperature in Fahrenheit: got %f, want %f", tempF, expectedTempF)
	}
}

func TestWeather_TempK(t *testing.T) {
	weather := &entity.Weather{TempC: 25.5}
	expectedTempK := float32(298.5)

	tempK := weather.TempK()

	if tempK != expectedTempK {
		t.Errorf("unexpected temperature in Kelvin: got %f, want %f", tempK, expectedTempK)
	}
}
