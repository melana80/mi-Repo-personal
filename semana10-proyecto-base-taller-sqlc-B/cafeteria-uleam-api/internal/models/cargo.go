package models

// Cargo representa el cargo o puesto de un empleado en la cafetería.
//
// La relación con Empleado es por ID: Empleado.CargoID apunta a Cargo.ID.
type Cargo struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Nombre      string `json:"nombre" gorm:"not null"`
	Descripcion string `json:"descripcion"`
}
