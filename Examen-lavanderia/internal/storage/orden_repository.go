// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import "github.com/joancema/examen-lavanderia/internal/models"

// OrdenRepository define el contrato de persistencia de Orden.
// Su implementación GORM (en orden_gorm.go) debe satisfacer EXACTAMENTE
// estas firmas. Observe que el repositorio NO contiene lógica de negocio:
// las reglas (validaciones, cálculo del total, anulación) viven en el service.
type OrdenRepository interface {
	Crear(a *models.Orden) error
	ObtenerPorID(id uint) (models.Orden, bool)
	Listar() ([]models.Orden, error)
	Actualizar(a *models.Orden) error
}
