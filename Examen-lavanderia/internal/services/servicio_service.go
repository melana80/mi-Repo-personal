// ARCHIVO BLOQUEADO — NO MODIFICAR
package services

import (
	"github.com/joancema/examen-lavanderia/internal/models"
	"github.com/joancema/examen-lavanderia/internal/storage"
)

// ServicioService contiene la lógica de negocio de la Entidad A.
// Está completo: úselo como ejemplo de cómo un service valida datos,
// devuelve errores de dominio y delega la persistencia al repository.
type ServicioService struct {
	repo storage.ServicioRepository
}

func NuevoServicioService(repo storage.ServicioRepository) *ServicioService {
	return &ServicioService{repo: repo}
}

func (s *ServicioService) Crear(h *models.Servicio) error {
	if h.Nombre == "" || h.PrecioUnitario <= 0 {
		return ErrDatosInvalidos
	}
	return s.repo.Crear(h)
}

func (s *ServicioService) ObtenerPorID(id uint) (models.Servicio, error) {
	h, ok := s.repo.ObtenerPorID(id)
	if !ok {
		return models.Servicio{}, ErrNoEncontrado
	}
	return h, nil
}

func (s *ServicioService) Listar() ([]models.Servicio, error) {
	return s.repo.Listar()
}
