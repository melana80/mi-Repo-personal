package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"


	"github.com/go-chi/chi/v5"

	"cafeteria-uleam-api/internal/models"
)

// ListarCategorias atiende GET /api/v1/categorias.
func (s *Server) ListarCategorias(w http.ResponseWriter, _ *http.Request) {
	categorias := s.Categorias.ListarC()
	RespondJSON(w, http.StatusOK, categorias)
}

// ObtenerCategoria atiende GET /api/v1/categorias/{id}.
func (s *Server) ObtenerCategoria(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	categoria, err := s.Categorias.ObtenerC(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	

	RespondJSON(w, http.StatusOK, categoria)
}

// CrearCategoria atiende POST /api/v1/categorias.
func (s *Server) CrearCategoria(w http.ResponseWriter, r *http.Request) {
	var nueva models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	creada, err := s.Categorias.CrearC(nueva)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creada)
}

// ActualizarCategoria atiende PUT /api/v1/categorias/{id}.
func (s *Server) ActualizarCategoria(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizada, err := s.Categorias.ActualizarC(id, datos)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, actualizada)
}

// BorrarCategoria atiende DELETE /api/v1/categorias/{id}.
func (s *Server) BorrarCategoria(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if err := s.Categorias.BorrarC(id); err != nil {
		RespondError(w, StatusDeError(err), err.Error())
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
