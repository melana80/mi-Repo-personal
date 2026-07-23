// ARCHIVO BLOQUEADO — NO MODIFICAR
package models

import "gorm.io/gorm"

// Servicio es la Entidad A: el catálogo de la lavandería.
// Este modelo está completo y sirve de plantilla para los que usted debe crear.
type Servicio struct {
	gorm.Model
	Nombre         string  `gorm:"size:120;not null" json:"nombre"`
	PrecioUnitario float64 `gorm:"not null" json:"precio_unitario"`
	Stock          uint    `gorm:"not null" json:"stock"`
	Activo         bool    `gorm:"not null;default:true" json:"activo"`
}
