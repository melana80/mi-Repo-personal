package storage

import (
	"gorm.io/gorm"

	"github.com/joancema/examen-lavanderia/internal/models"
)

type OrdenGORM struct {
	db *gorm.DB
}

func NuevaOrdenGORM(db *gorm.DB) *OrdenGORM {
	return &OrdenGORM{db: db}
}

func (r *OrdenGORM) Crear(a *models.Orden) error {
	return r.db.Create(a).Error
}

func (r *OrdenGORM) ObtenerPorID(id uint) (models.Orden, bool) {
	var a models.Orden
	if err := r.db.First(&a, id).Error; err != nil {
		return models.Orden{}, false
	}
	return a, true
}

func (r *OrdenGORM) Listar() ([]models.Orden, error) {
	var lista []models.Orden
	err := r.db.Find(&lista).Error
	return lista, err
}

func (r *OrdenGORM) Actualizar(a *models.Orden) error {
	return r.db.Save(a).Error
}
