// ARCHIVO BLOQUEADO — NO MODIFICAR
package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NuevoRouter registra todas las rutas de la API. Este archivo es el
// contrato HTTP del examen: los tests httptest de acceptance/ atacan
// exactamente estas rutas.
func NuevoRouter(
	servicios *ServicioHandler,
	clientes *ClienteHandler,
	ordenes *OrdenHandler,
) http.Handler {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/servicios", func(r chi.Router) {
			r.Get("/", servicios.Listar)
			r.Post("/", servicios.Crear)
		})

		r.Route("/clientes", func(r chi.Router) {
			r.Get("/", clientes.Listar)
			r.Post("/", clientes.Crear)
			r.Get("/{id}", clientes.ObtenerPorID)
		})

		r.Route("/ordenes", func(r chi.Router) {
			r.Get("/", ordenes.Listar)
			r.Post("/", ordenes.Crear)
			r.Get("/{id}", ordenes.ObtenerPorID)
			r.Post("/{id}/cancelar", ordenes.Cancelar)
		})
	})

	return r
}
