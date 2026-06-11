// Package storage gestiona el almacenamiento en memoria de la cafetería.
//
// El tipo Memoria mantiene en un solo lugar todos los datos del dominio:
// Empleados y Cargos.
package storage

import (
	"sync"

	"cafeteria-uleam-api/internal/models"
)

// Memoria es un almacén unificado de la cafetería.
type Memoria struct {
	empleados      []models.Empleado
	nextEmpleadoID int

	cargos      []models.Cargo
	nextCargoID int

	mu sync.Mutex
}

// NuevaMemoria crea un almacén vacío y listo para usar.
func NuevaMemoria() *Memoria {
	return &Memoria{
		empleados:      []models.Empleado{},
		nextEmpleadoID: 1,
		cargos:         []models.Cargo{},
		nextCargoID:    1,
	}
}

// =========================================================
// EMPLEADOS
// =========================================================

// SeedEmpleados carga empleados iniciales en memoria.
func (m *Memoria) SeedEmpleados() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.empleados = []models.Empleado{
		{ID: 1, Nombre: "Ana", Apellido: "Torres", Salario: 520.00, HorasSemana: 40, CargoID: 1},
		{ID: 2, Nombre: "Luis", Apellido: "Mero", Salario: 480.50, HorasSemana: 36, CargoID: 1},
		{ID: 3, Nombre: "Carla", Apellido: "Vera", Salario: 500.00, HorasSemana: 40, CargoID: 2},
		{ID: 4, Nombre: "Jorge", Apellido: "Pin", Salario: 610.75, HorasSemana: 40, CargoID: 3},
		{ID: 5, Nombre: "María", Apellido: "Loor", Salario: 590.00, HorasSemana: 32, CargoID: 3},
		{ID: 6, Nombre: "Pedro", Apellido: "Cedeño", Salario: 900.00, HorasSemana: 40, CargoID: 4},
	}
	m.nextEmpleadoID = 7
}

// ListarEmpleados devuelve todos los empleados en memoria.
func (m *Memoria) ListarEmpleados() []models.Empleado {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Empleado, len(m.empleados))
	copy(copia, m.empleados)
	return copia
}

// BuscarEmpleadoPorID devuelve el empleado con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarEmpleadoPorID(id int) (models.Empleado, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, p := range m.empleados {
		if p.ID == id {
			return p, true
		}
	}
	return models.Empleado{}, false
}

// CrearEmpleado agrega un empleado nuevo y devuelve el empleado con ID asignado.
func (m *Memoria) CrearEmpleado(p models.Empleado) models.Empleado {
	m.mu.Lock()
	defer m.mu.Unlock()

	p.ID = m.nextEmpleadoID
	m.nextEmpleadoID++
	m.empleados = append(m.empleados, p)
	return p
}

// ActualizarEmpleado reemplaza el empleado con el ID dado.
func (m *Memoria) ActualizarEmpleado(id int, datos models.Empleado) (models.Empleado, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.empleados {
		if p.ID == id {
			datos.ID = id
			m.empleados[i] = datos
			return datos, true
		}
	}
	return models.Empleado{}, false
}

// BorrarEmpleado elimina el empleado con el ID dado.
func (m *Memoria) BorrarEmpleado(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, p := range m.empleados {
		if p.ID == id {
			m.empleados = append(m.empleados[:i], m.empleados[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// CARGOS
// =========================================================

// SeedCargos carga cargos iniciales que coinciden con CargoID de los empleados pre-cargados.
func (m *Memoria) SeedCargos() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cargos = []models.Cargo{
		{ID: 1, Nombre: "Barista", Descripcion: "Prepara y sirve cafés y bebidas calientes"},
		{ID: 2, Nombre: "Cajero", Descripcion: "Atiende la caja y cobra los pedidos"},
		{ID: 3, Nombre: "Cocina", Descripcion: "Prepara alimentos y controla insumos"},
		{ID: 4, Nombre: "Administración", Descripcion: "Gestiona personal, compras y reportes"},
	}
	m.nextCargoID = 5
}

// ListarCargos devuelve todas las cargos en memoria.
func (m *Memoria) ListarCargos() []models.Cargo {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Cargo, len(m.cargos))
	copy(copia, m.cargos)
	return copia
}

// BuscarCargoPorID devuelve la cargo con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarCargoPorID(id int) (models.Cargo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, c := range m.cargos {
		if c.ID == id {
			return c, true
		}
	}
	return models.Cargo{}, false
}

// CrearCargo agrega una cargo nueva y devuelve la cargo con ID asignado.
func (m *Memoria) CrearCargo(c models.Cargo) models.Cargo {
	m.mu.Lock()
	defer m.mu.Unlock()

	c.ID = m.nextCargoID
	m.nextCargoID++
	m.cargos = append(m.cargos, c)
	return c
}

// ActualizarCargo reemplaza la cargo con el ID dado.
func (m *Memoria) ActualizarCargo(id int, datos models.Cargo) (models.Cargo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.cargos {
		if c.ID == id {
			datos.ID = id
			m.cargos[i] = datos
			return datos, true
		}
	}
	return models.Cargo{}, false
}

// BorrarCargo elimina la cargo con el ID dado.
func (m *Memoria) BorrarCargo(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.cargos {
		if c.ID == id {
			m.cargos = append(m.cargos[:i], m.cargos[i+1:]...)
			return true
		}
	}
	return false
}
