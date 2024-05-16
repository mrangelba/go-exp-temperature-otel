package weatherapi_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mrangelba/go-exp-temperature/service-b/internal/infra/gateway/weatherapi"
)

func TestWeatherGateway_Get(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the query parameters
		q := r.URL.Query()
		if q.Get("key") != "your_weather_api_key" {
			t.Errorf("unexpected API key: %s", q.Get("key"))
		}
		if q.Get("q") != "your_city_name" {
			t.Errorf("unexpected city name: %s", q.Get("q"))
		}
		if q.Get("aqi") != "no" {
			t.Errorf("unexpected AQI value: %s", q.Get("aqi"))
		}

		// Send a mock response
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"current": {"temp_c": 25.5}}`))
	}))
	defer server.Close()

	weatherapi.BASE_URL = server.URL

	// Create a new weather gateway with the mock server
	gateway := weatherapi.NewWeatherGateway(server.Client(), "your_weather_api_key")

	// Call the Get method
	ctx := context.Background()
	weather, err := gateway.Get(ctx, "your_city_name")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify the temperature value
	expectedTempC := float32(25.5)
	if weather.TempC != expectedTempC {
		t.Errorf("unexpected temperature value: %f, expected: %f", weather.TempC, expectedTempC)
	}
}

func TestWeatherGatewayDecodeError_Get(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the query parameters
		q := r.URL.Query()
		if q.Get("key") != "your_weather_api_key" {
			t.Errorf("unexpected API key: %s", q.Get("key"))
		}
		if q.Get("q") != "your_city_name" {
			t.Errorf("unexpected city name: %s", q.Get("q"))
		}
		if q.Get("aqi") != "no" {
			t.Errorf("unexpected AQI value: %s", q.Get("aqi"))
		}

		// Send a mock response
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"current": {"temp_c": 25.5}`))
	}))
	defer server.Close()

	weatherapi.BASE_URL = server.URL

	// Create a new weather gateway with the mock server
	gateway := weatherapi.NewWeatherGateway(server.Client(), "your_weather_api_key")

	// Call the Get method
	ctx := context.Background()
	_, err := gateway.Get(ctx, "your_city_name")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestWeatherGatewayServerError_Get(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	weatherapi.BASE_URL = server.URL

	// Create a new weather gateway with the mock server
	gateway := weatherapi.NewWeatherGateway(server.Client(), "your_weather_api_key")

	// Call the Get method
	ctx := context.Background()
	_, err := gateway.Get(ctx, "your_city_name")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestWeatherGatewayRequestError_Get(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer server.Close()

	weatherapi.BASE_URL = server.URL

	// Create a new weather gateway with the mock server
	gateway := weatherapi.NewWeatherGateway(server.Client(), "your_weather_api_key")

	// Call the Get method
	//lint:ignore SA1012 this is a test
	_, err := gateway.Get(nil, "your_city_name")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestWeatherGatewayURLInvalid_Get(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer server.Close()

	weatherapi.BASE_URL = "http://"

	// Create a new weather gateway with the mock server
	gateway := weatherapi.NewWeatherGateway(server.Client(), "your_weather_api_key")

	// Call the Get method
	//lint:ignore SA1012 this is a test
	_, err := gateway.Get(nil, "your_city_name")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
