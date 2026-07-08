package handlers

import (
    "cafeteria-uleam-api/internal/service"
)
// Server agrupa las dependencias compartidas por los handlers.
type Server struct {
    Productos     *service.ProductoService
    Categorias    *service.CategoriaService
    Autenticacion *service.AutenticacionService
}

func NewServer(productos *service.ProductoService, categorias *service.CategoriaService, autenticacion *service.AutenticacionService) *Server {
    return &Server{
		Productos:     productos,
		Categorias:    categorias,
		Autenticacion: autenticacion,
    }
}
