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

type OrdenHandler struct {
	servicio *services.OrdenService
}

func NuevaOrdenHandler(s *services.OrdenService) *OrdenHandler {
	return &OrdenHandler{servicio: s}
}

func (h *OrdenHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var req crearOrdenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}
	if req.ServicioID == 0 || req.ClienteID == 0 || req.Cantidad == 0 {
		RespondError(w, http.StatusUnprocessableEntity, services.ErrDatosInvalidos.Error())
		return
	}
	orden := models.Orden{ServicioID: req.ServicioID, ClienteID: req.ClienteID, Cantidad: req.Cantidad}
	if err := h.servicio.Crear(&orden); err != nil {
		switch {
		default:
			RespondError(w, http.StatusInternalServerError, err.Error())
		case errors.Is(err, services.ErrDatosInvalidos), errors.Is(err, services.ErrReferenciaInvalida):
			RespondError(w, http.StatusUnprocessableEntity, err.Error())
		case errors.Is(err, services.ErrStockInsuficiente), errors.Is(err, services.ErrEstadoInvalido):
			RespondError(w, http.StatusConflict, err.Error())
		}
		return
	}
	RespondJSON(w, http.StatusCreated, orden)
}

func (h *OrdenHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.servicio.Listar()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, lista)
}

func (h *OrdenHandler) ObtenerPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	a, err := h.servicio.ObtenerPorID(uint(id))
	if err != nil {
		if errors.Is(err, services.ErrNoEncontrado) {
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, a)
}

func (h *OrdenHandler) Cancelar(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	if err := h.servicio.Cancelar(uint(id)); err != nil {
		switch {
		default:
			RespondError(w, http.StatusInternalServerError, err.Error())
		case errors.Is(err, services.ErrNoEncontrado):
			RespondError(w, http.StatusNotFound, err.Error())
		case errors.Is(err, services.ErrEstadoInvalido):
			RespondError(w, http.StatusConflict, err.Error())
		}
		return
	}
	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "orden cancelada"})
}

type crearOrdenRequest struct {
	ServicioID uint `json:"servicio_id"`
	ClienteID  uint `json:"cliente_id"`
	Cantidad   uint `json:"cantidad"`
}
