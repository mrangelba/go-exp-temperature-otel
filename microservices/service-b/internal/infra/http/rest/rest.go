package rest

import (
	"log"
	"net/http"

	"github.com/mrangelba/go-exp-temperature/service-b/internal/di"
	"github.com/riandyrn/otelchi"

	"github.com/go-chi/chi/v5"
)

func Initialize() {
	webController := di.ConfigWebController()

	r := chi.NewRouter()
	r.Use(otelchi.Middleware("service-b", otelchi.WithChiRoutes(r)))
	r.Get("/cep/{cep}", webController.Get)

	log.Println("Server running on port 8081")
	http.ListenAndServe(":8081", r)
}
