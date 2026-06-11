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
	Storage *storage.Memoria
}

// NewServer construye un Server listo para usar.
func NewServer(s *storage.Memoria) *Server {
	return &Server{Storage: s}
}

// ListarProductos atiende GET /api/v1/productos.
func (s *Server) ListarProductos(w http.ResponseWriter, _ *http.Request) {
	productos := s.Storage.ListarProductos()


	RespondJson(w, http.StatusOK, productos)
}


// ObtenerProducto atiende GET /api/v1/productos/{id}.
func (s *Server) ObtenerProducto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	producto, encontrado := s.Storage.BuscarProductoPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "producto no encontrado")
		return
	}

	//reemplazar con respondJson

	RespondJson(w, http.StatusOK, producto)
	

}


// CrearProducto atiende POST /api/v1/productos.
func (s *Server) CrearProducto(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Producto
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nuevo.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if nuevo.Precio < 0 {
		http.Error(w, "el precio no puede ser negativo", http.StatusBadRequest)
		return
	}

	creado := s.Storage.CrearProducto(nuevo)

	RespondJson(w, http.StatusCreated, creado)
	
}

// ActualizarProducto atiende PUT /api/v1/productos/{id}.
func (s *Server) ActualizarProducto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	var datos models.Producto
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.Nombre) == "" {
		http.Error(w, "el campo nombre es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.Precio < 0 {
		http.Error(w, "el precio no puede ser negativo", http.StatusBadRequest)
		return
	}

	actualizado, encontrado := s.Storage.ActualizarProducto(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "producto no encontrado")
		return
	}

	RespondJson(w, http.StatusOK, actualizado)
	
}

// BorrarProducto atiende DELETE /api/v1/productos/{id}.
func (s *Server) BorrarProducto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	if !s.Storage.BorrarProducto(id) {
		RespondError(w, http.StatusNotFound, "producto no encontrado")
		return
	}

	
	RespondJson(w, http.StatusNoContent, nil)
}

// TODO TALLER: agregar handlers para Categoria como métodos de Server.
// Deben seguir el mismo patrón que los de Producto:
// - ListarCategorias
// - ObtenerCategoria
// - CrearCategoria
// - ActualizarCategoria
// - BorrarCategoria
