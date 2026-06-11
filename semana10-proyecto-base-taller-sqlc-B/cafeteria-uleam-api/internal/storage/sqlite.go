package storage

import (
	"gorm.io/gorm"

	"cafeteria-uleam-api/internal/models"
)

// AlmacenSQLite implementa la interfaz Almacen usando GORM sobre SQLite.
//
// Fíjense: los métodos tienen EXACTAMENTE las mismas firmas que los de Memoria.
// Por eso el Server y los handlers no se enteran de cuál de los dos reciben.
type AlmacenSQLite struct {
	db *gorm.DB
}

// NuevoAlmacenSQLite envuelve una conexión *gorm.DB ya abierta.
func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

// =========================================================
// EMPLEADOS
// =========================================================

func (a *AlmacenSQLite) ListarEmpleados() []models.Empleado {
	var empleados []models.Empleado
	a.db.Find(&empleados)
	return empleados
}

func (a *AlmacenSQLite) BuscarEmpleadoPorID(id int) (models.Empleado, bool) {
	var p models.Empleado
	if err := a.db.First(&p, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return models.Empleado{}, false
	}
	return p, true
}

func (a *AlmacenSQLite) CrearEmpleado(p models.Empleado) models.Empleado {
	a.db.Create(&p) // GORM rellena el ID autogenerado en &p
	return p
}

func (a *AlmacenSQLite) ActualizarEmpleado(id int, datos models.Empleado) (models.Empleado, bool) {
	var existente models.Empleado
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Empleado{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarEmpleado(id int) bool {
	res := a.db.Delete(&models.Empleado{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// CARGOS
// =========================================================

func (a *AlmacenSQLite) ListarCargos() []models.Cargo {
	var cargos []models.Cargo
	a.db.Find(&cargos)
	return cargos
}

func (a *AlmacenSQLite) BuscarCargoPorID(id int) (models.Cargo, bool) {
	var c models.Cargo
	if err := a.db.First(&c, id).Error; err != nil {
		return models.Cargo{}, false
	}
	return c, true
}

func (a *AlmacenSQLite) CrearCargo(c models.Cargo) models.Cargo {
	a.db.Create(&c)
	return c
}

func (a *AlmacenSQLite) ActualizarCargo(id int, datos models.Cargo) (models.Cargo, bool) {
	var existente models.Cargo
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Cargo{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarCargo(id int) bool {
	res := a.db.Delete(&models.Cargo{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// SEEDS
// =========================================================

// SembrarSiVacio inserta datos iniciales solo si aún no hay cargos.
// Así no duplicamos datos en cada arranque del servidor.
func (a *AlmacenSQLite) SembrarSiVacio() {
	var n int64
	a.db.Model(&models.Cargo{}).Count(&n)
	if n > 0 {
		return
	}

	cargos := []models.Cargo{
		{ID: 1, Nombre: "Barista", Descripcion: "Prepara y sirve cafés y bebidas calientes"},
		{ID: 2, Nombre: "Cajero", Descripcion: "Atiende la caja y cobra los pedidos"},
		{ID: 3, Nombre: "Cocina", Descripcion: "Prepara alimentos y controla insumos"},
		{ID: 4, Nombre: "Administración", Descripcion: "Gestiona personal, compras y reportes"},
	}
	a.db.Create(&cargos)

	empleados := []models.Empleado{
		{ID: 1, Nombre: "Ana", Apellido: "Torres", Salario: 520.00, HorasSemana: 40, CargoID: 1},
		{ID: 2, Nombre: "Luis", Apellido: "Mero", Salario: 480.50, HorasSemana: 36, CargoID: 1},
		{ID: 3, Nombre: "Carla", Apellido: "Vera", Salario: 500.00, HorasSemana: 40, CargoID: 2},
		{ID: 4, Nombre: "Jorge", Apellido: "Pin", Salario: 610.75, HorasSemana: 40, CargoID: 3},
		{ID: 5, Nombre: "María", Apellido: "Loor", Salario: 590.00, HorasSemana: 32, CargoID: 3},
		{ID: 6, Nombre: "Pedro", Apellido: "Cedeño", Salario: 900.00, HorasSemana: 40, CargoID: 4},
	}
	a.db.Create(&empleados)
}

// Chequeo en tiempo de compilación: AlmacenSQLite debe cumplir Almacen.
var _ Almacen = (*AlmacenSQLite)(nil)
