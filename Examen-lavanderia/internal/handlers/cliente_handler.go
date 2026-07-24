package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/joancema/examen-lavanderia/internal/models"
	"github.com/joancema/examen-lavanderia/internal/services"
)


type ClienteHandler struct {
	servicio *services.ClienteService
}

func NuevoClienteHandler(s *services.ClienteService) *ClienteHandler {
	return &ClienteHandler{servicio: s}
}

func (h *ClienteHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var req crearClienteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}
	if req.Nombre == "" || req.Cedula == "" || req.Telefono == "" {
		RespondError(w, http.StatusUnprocessableEntity, services.ErrDatosInvalidos.Error())
		return
	}
	cliente := models.Cliente{Nombre: req.Nombre, Cedula: req.Cedula, Telefono: req.Telefono}
	if err := h.servicio.Crear(&cliente); err != nil {
		switch {
		default:
			RespondError(w, http.StatusInternalServerError, err.Error())
		case errors.Is(err, services.ErrDatosInvalidos):
			RespondError(w, http.StatusUnprocessableEntity, err.Error())
		}
		return
	}
	RespondJSON(w, http.StatusCreated, cliente)
}

func (h *ClienteHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.servicio.Listar()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, lista)
}

func (h *ClienteHandler) ObtenerPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	c, err := h.servicio.ObtenerPorID(uint(id))
	if err != nil {
		if errors.Is(err, services.ErrNoEncontrado) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, c)
}

type crearClienteRequest struct {
	Nombre   string `json:"nombre"`
	Cedula   string `json:"cedula"`
	Telefono string `json:"telefono"`
}
