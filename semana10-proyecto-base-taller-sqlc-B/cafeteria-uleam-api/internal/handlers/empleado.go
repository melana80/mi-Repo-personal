// Package handlers contiene los handlers HTTP de la API de cafetería.
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"cafeteria-uleam-api/internal/models"
	"cafeteria-uleam-api/internal/storage"
)

// Server agrupa las dependencias compartidas por los handlers.
// Recibe el storage por inyección de dependencias desde main.
type Server struct {
	Storage storage.Almacen
}

// NewServer construye un Server listo para usar.
func NewServer(s storage.Almacen) *Server {
	return &Server{Storage: s}
}

// ListarEmpleados atiende GET /api/v1/empleados.
func (s *Server) ListarEmpleados(w http.ResponseWriter, _ *http.Request) {
	empleados := s.Storage.ListarEmpleados()
	RespondJSON(w, http.StatusOK, empleados)
}

// ObtenerEmpleado atiende GET /api/v1/empleados/{id}.
func (s *Server) ObtenerEmpleado(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	empleado, encontrado := s.Storage.BuscarEmpleadoPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "empleado no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, empleado)
}

// CrearEmpleado atiende POST /api/v1/empleados.
func (s *Server) CrearEmpleado(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Empleado
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nuevo.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if nuevo.Salario < 0 {
		RespondError(w, http.StatusBadRequest, "el salario no puede ser negativo")
		return
	}

	creado := s.Storage.CrearEmpleado(nuevo)
	RespondJSON(w, http.StatusCreated, creado)
}

// ActualizarEmpleado atiende PUT /api/v1/empleados/{id}.
func (s *Server) ActualizarEmpleado(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos models.Empleado
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(datos.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if datos.Salario < 0 {
		RespondError(w, http.StatusBadRequest, "el salario no puede ser negativo")
		return
	}

	actualizado, encontrado := s.Storage.ActualizarEmpleado(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "empleado no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

// BorrarEmpleado atiende DELETE /api/v1/empleados/{id}.
func (s *Server) BorrarEmpleado(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if !s.Storage.BorrarEmpleado(id) {
		RespondError(w, http.StatusNotFound, "empleado no encontrado")
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
