package service

import (

	"strings"

	"cafeteria-uleam-api/internal/models"
	"cafeteria-uleam-api/internal/storage"
)
// Producto representa un producto de la cafetería.
type CategoriaService struct {
	repo storage.CategoriaRepository
}

// CrearProducto crea un nuevo producto.
func NuevaCategoriaService(repo storage.CategoriaRepository) *CategoriaService {
	return &CategoriaService{repo: repo}
}

//listar
func (s *CategoriaService) ListarC() []models.Categoria {
	return s.repo.ListarCategorias()
}
//obtener
func (s *CategoriaService) ObtenerC(id int) (models.Categoria, error) {
	c, ok := s.repo.BuscarCategoriaPorID(id)
	if !ok {
		return models.Categoria{}, ErrNoEncontrado
	}
	return c, nil
}

//validar
func ValidarCategoria(c models.Categoria) error {
	if strings.TrimSpace(c.Nombre) == "" {
		return ErrNombreVacio
	}
	return nil
}

//crear
func (s *CategoriaService) CrearC(c models.Categoria) (models.Categoria, error) {
	if err := ValidarCategoria(c); err != nil {
		return models.Categoria{}, err
	}
	return s.repo.CrearCategoria(c), nil
}

//actualizar
func (s *CategoriaService) ActualizarC(id int, c models.Categoria) (models.Categoria, error) {
	if err := ValidarCategoria(c); err != nil {
		return models.Categoria{}, err
	}
	actualizado, ok := s.repo.ActualizarCategoria(id, c)
	if !ok {
		return models.Categoria{}, ErrNoEncontrado
	}
	return actualizado, nil
}

//borrar
func (s *CategoriaService) BorrarC(id int) error {
	if !s.repo.BorrarCategoria(id) {
		return ErrNoEncontrado
	}
	return nil
}