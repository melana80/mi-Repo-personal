package models

import "gorm.io/gorm"

type Orden struct {
	gorm.Model
	ServicioID uint    `gorm:"not null" json:"servicio_id"`
	ClienteID  uint    `gorm:"not null" json:"cliente_id"`
	Cantidad   uint    `gorm:"not null" json:"cantidad"`
	Estado     string  `gorm:"not null" json:"estado"`
	Total      float64 `gorm:"not null" json:"total"`
}
