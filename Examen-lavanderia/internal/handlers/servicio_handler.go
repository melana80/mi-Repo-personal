// ARCHIVO BLOQUEADO — NO MODIFICAR
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/joancema/examen-lavanderia/internal/models"
	"github.com/joancema/examen-lavanderia/internal/services"
)

// ServicioHandler expone la Entidad A por HTTP.
// Está completo: observe cómo decodifica el body, llama al service y
// MAPEA los errores de dominio a status codes. Ese mapeo es exactamente
// lo que usted debe replicar en sus propios handlers.
type ServicioHandler struct {
	servicio *services.ServicioService
}

func NuevoServicioHandler(s *services.ServicioService) *ServicioHandler {
	return &ServicioHandler{servicio: s}
}

func (h *ServicioHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.servicio.Listar()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, lista)
}

func (h *ServicioHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var servicio models.Servicio
	if err := json.NewDecoder(r.Body).Decode(&servicio); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}
	if err := h.servicio.Crear(&servicio); err != nil {
		switch {
		case errors.Is(err, services.ErrDatosInvalidos):
			RespondError(w, http.StatusUnprocessableEntity, err.Error())
		default:
			RespondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	RespondJSON(w, http.StatusCreated, servicio)
}
