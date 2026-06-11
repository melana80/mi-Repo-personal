// Package models define las entidades del dominio cafetería.
package models

// Empleado representa a un empleado de la cafetería.
//
// CargoID referencia el ID de un Cargo por número (foreign key).
// Decisión arquitectónica: usamos ID en lugar de struct anidado. GORM usa
// ese mismo ID como clave foránea, así que el modelo casi no cambia.
type Empleado struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Nombre      string  `json:"nombre" gorm:"not null"`
	Apellido    string  `json:"apellido" gorm:"not null"`
	Salario     float64 `json:"salario"`
	HorasSemana int     `json:"horas_semana"`
	CargoID     int     `json:"cargo_id"`
}
