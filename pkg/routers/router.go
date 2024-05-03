package routers

import (
	"net/http"

	"github.com/oxxi/accel-one/internal/api"
	"github.com/oxxi/accel-one/internal/service"
)

var RegisterRouter = func(router *http.ServeMux) {
	service := service.NewContractService()
	handlers := api.NewContactHandler(service)

	router.HandleFunc("POST /contacts", handlers.SaveContact)
	router.HandleFunc("GET /contacts/{id}", handlers.GetContactById)
	router.HandleFunc("DELETE /contacts/{id}", handlers.DeleteContact)
	router.HandleFunc("PUT /contacts/{id}", handlers.UpdateContact)
}
