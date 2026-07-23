package models

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model
	Nombre    string `gorm:"size:120;not null" json:"nombre"`
	Cedula    string `gorm:"size:20;not null;uniqueIndex" json:"cedula"`
	Telefono  string `gorm:"size:20;not null" json:"telefono"`
}
