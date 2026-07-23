// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import "github.com/joancema/examen-lavanderia/internal/models"

// ServicioRepository define el contrato de persistencia de la Entidad A.
type ServicioRepository interface {
	Crear(h *models.Servicio) error
	ObtenerPorID(id uint) (models.Servicio, bool)
	Listar() ([]models.Servicio, error)
	Actualizar(h *models.Servicio) error
}
