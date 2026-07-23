// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"gorm.io/gorm"

	"github.com/joancema/examen-lavanderia/internal/models"
)

// ServicioGORM implementa ServicioRepository sobre GORM.
// Esta implementación está completa: úsela como plantilla para ClienteGORM
// y OrdenGORM, que usted debe implementar.
type ServicioGORM struct {
	db *gorm.DB
}

func NuevoServicioGORM(db *gorm.DB) *ServicioGORM {
	return &ServicioGORM{db: db}
}

func (r *ServicioGORM) Crear(h *models.Servicio) error {
	return r.db.Create(h).Error
}

func (r *ServicioGORM) ObtenerPorID(id uint) (models.Servicio, bool) {
	var h models.Servicio
	if err := r.db.First(&h, id).Error; err != nil {
		return models.Servicio{}, false
	}
	return h, true
}

func (r *ServicioGORM) Listar() ([]models.Servicio, error) {
	var lista []models.Servicio
	err := r.db.Find(&lista).Error
	return lista, err
}

func (r *ServicioGORM) Actualizar(h *models.Servicio) error {
	return r.db.Save(h).Error
}
