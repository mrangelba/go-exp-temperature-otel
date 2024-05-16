package viacep_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mrangelba/go-exp-temperature/service-b/internal/domain/entity"
	"github.com/mrangelba/go-exp-temperature/service-b/internal/infra/gateway/viacep"
	"github.com/mrangelba/go-exp-temperature/service-b/pkg/error_handle"
)

func TestValidCEP_Get(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"cep":"96200-006","localidade":"City","erro":false}`))
	}))
	defer server.Close()
	viacep.BASE_URL = server.URL

	// Create a new CEP gateway with the test server's client
	gateway := viacep.NewCepGateway(server.Client())

	// Call the Get method with a valid CEP
	cep := "96200-006"
	result, err := gateway.Get(context.Background(), cep)

	// Check if there was an error
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the result
	expected := &entity.CEP{
		Cep:      "96200-006",
		CityName: "City",
	}
	if result.Cep != expected.Cep || result.CityName != expected.CityName {
		t.Errorf("Expected result %v, got %v", expected, result)
	}
}

func TestInvalidCEP_Get(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"cep":"abcde","localidade":"","erro":true}`))
	}))
	defer server.Close()
	viacep.BASE_URL = server.URL

	// Create a new CEP gateway with the test server's client
	gateway := viacep.NewCepGateway(server.Client())

	// Call the Get method with an invalid CEP
	cep := "abcde"
	_, err := gateway.Get(context.Background(), cep)

	// Check if the error is ErrNotFound
	if err != error_handle.ErrNotFound {
		t.Errorf("Expected error %v, got %v", error_handle.ErrNotFound, err)
	}
}

func TestServerError_Get(t *testing.T) {
	// Create a test server that returns a server error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	viacep.BASE_URL = server.URL

	// Create a new CEP gateway with the test server's client
	gateway := viacep.NewCepGateway(server.Client())

	// Call the Get method
	cep := "12345-678"
	_, err := gateway.Get(context.Background(), cep)

	// Check if the error is ErrUnprocessableEntity
	if err != error_handle.ErrUnprocessableEntity {
		t.Errorf("Expected error %v, got %v", error_handle.ErrUnprocessableEntity, err)
	}
}

func TestServerRequestError_Get(t *testing.T) {
	// Create a test server that returns a server error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer server.Close()
	viacep.BASE_URL = server.URL

	// Create a new CEP gateway with the test server's client
	gateway := viacep.NewCepGateway(server.Client())

	// Call the Get method
	//lint:ignore SA1012 this is a test
	_, err := gateway.Get(nil, "")

	// Check if the error is ErrUnprocessableEntity
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestValidCEPDecodeError_Get(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"cep":"96200-006","localidade":"City","erro":false`))
	}))
	defer server.Close()
	viacep.BASE_URL = server.URL

	// Create a new CEP gateway with the test server's client
	gateway := viacep.NewCepGateway(server.Client())

	// Call the Get method with a valid CEP
	cep := "96200-006"
	_, err := gateway.Get(context.Background(), cep)

	// Check if there was an error
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
