package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"cafeteria-uleam-api/internal/models"
)

// ListarCargos atiende GET /api/v1/cargos.
func (s *Server) ListarCargos(w http.ResponseWriter, _ *http.Request) {
	cargos := s.Storage.ListarCargos()
	RespondJSON(w, http.StatusOK, cargos)
}

// ObtenerCargo atiende GET /api/v1/cargos/{id}.
func (s *Server) ObtenerCargo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	cargo, encontrado := s.Storage.BuscarCargoPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "cargo no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, cargo)
}

// CrearCargo atiende POST /api/v1/cargos.
func (s *Server) CrearCargo(w http.ResponseWriter, r *http.Request) {
	var nueva models.Cargo
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nueva.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}

	creada := s.Storage.CrearCargo(nueva)
	RespondJSON(w, http.StatusCreated, creada)
}

// ActualizarCargo atiende PUT /api/v1/cargos/{id}.
func (s *Server) ActualizarCargo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos models.Cargo
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(datos.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}

	actualizada, encontrada := s.Storage.ActualizarCargo(id, datos)
	if !encontrada {
		RespondError(w, http.StatusNotFound, "cargo no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, actualizada)
}

// BorrarCargo atiende DELETE /api/v1/cargos/{id}.
func (s *Server) BorrarCargo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if !s.Storage.BorrarCargo(id) {
		RespondError(w, http.StatusNotFound, "cargo no encontrado")
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
