package services

import (
	"github.com/joancema/examen-lavanderia/internal/models"
	"github.com/joancema/examen-lavanderia/internal/storage"
)

type ClienteService struct {
	repo storage.ClienteRepository
}

func NuevoClienteService(repo storage.ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (s *ClienteService) Crear(c *models.Cliente) error {
	if c.Nombre == "" || c.Cedula == "" || c.Telefono == "" {
		return ErrDatosInvalidos
	}
	return s.repo.Crear(c)
}

func (s *ClienteService) ObtenerPorID(id uint) (models.Cliente, error) {
	c, ok := s.repo.ObtenerPorID(id)
	if !ok {
		return models.Cliente{}, ErrNoEncontrado
	}
	return c, nil
}

func (s *ClienteService) Listar() ([]models.Cliente, error) {
	return s.repo.Listar()
}
