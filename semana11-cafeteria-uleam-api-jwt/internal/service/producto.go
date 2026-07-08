package service

import (
	"strings"
	"cafeteria-uleam-api/internal/models"
	"cafeteria-uleam-api/internal/storage"
	

)

// Producto representa un producto de la cafetería.
type ProductoService struct {
	repo storage.ProductoRepository 
}

// CrearProducto crea un nuevo producto.
func NuevaProductoService(repo storage.ProductoRepository) *ProductoService {
	return &ProductoService{repo: repo} 
}

func (s *ProductoService) Listar() []models.Producto {
	return s.repo.ListarProductos()
}

func (s *ProductoService) Obtener(id int) (models.Producto, error) {
	p, ok := s.repo.BuscarProductoPorID(id);
	if !ok {
		return models.Producto{}, ErrNoEncontrado
	}
	return p, nil
}
//validar producto
func ValidarProducto(p models.Producto) error {
	if strings.TrimSpace(p.Nombre) == "" {
		return ErrNombreVacio
	}
	if p.Precio <= 0 {
		return ErrPrecioNegativo
	}
	return nil
}

func (s *ProductoService) CrearP(p models.Producto) (models.Producto, error) {
	if err := ValidarProducto(p); err != nil {
		return models.Producto{}, err
	}
	return s.repo.CrearProducto(p), nil

}

func (s *ProductoService) ActualizarP(id int, p models.Producto) (models.Producto, error) {
	if err := ValidarProducto(p); err != nil {
		return models.Producto{}, err
	}
	actualizado, ok := s.repo.ActualizarProducto(id, p)
	if !ok {
		return models.Producto{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *ProductoService) BorrarP(id int) error {
	if !s.repo.BorrarProducto(id) {
		return ErrNoEncontrado
	}
	return nil
}