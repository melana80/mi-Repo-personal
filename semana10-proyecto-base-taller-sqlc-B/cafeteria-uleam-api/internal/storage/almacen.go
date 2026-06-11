package storage

import "cafeteria-uleam-api/internal/models"

// Almacen define QUÉ sabe hacer un almacén de la cafetería, sin decir CÓMO.
//
// Memoria (slices) ya cumple esta interfaz sin cambios — por el duck typing
// que vimos en S3 — y AlmacenSQLite (GORM) la cumple igual. El Server depende
// de esta interfaz, no de una implementación concreta: por eso podemos cambiar
// el backend de almacenamiento sin tocar un solo handler.
type Almacen interface {
	// Empleados
	ListarEmpleados() []models.Empleado
	BuscarEmpleadoPorID(id int) (models.Empleado, bool)
	CrearEmpleado(p models.Empleado) models.Empleado
	ActualizarEmpleado(id int, datos models.Empleado) (models.Empleado, bool)
	BorrarEmpleado(id int) bool

	// Cargos
	ListarCargos() []models.Cargo
	BuscarCargoPorID(id int) (models.Cargo, bool)
	CrearCargo(c models.Cargo) models.Cargo
	ActualizarCargo(id int, datos models.Cargo) (models.Cargo, bool)
	BorrarCargo(id int) bool
}

// Chequeo en tiempo de compilación: si Memoria dejara de cumplir Almacen,
// el proyecto NO compila. Red de seguridad opcional.
var _ Almacen = (*Memoria)(nil)
