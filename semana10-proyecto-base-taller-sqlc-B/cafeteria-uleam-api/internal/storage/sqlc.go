package storage
import (
	"context"
	"database/sql"
	"cafeteria-uleam-api/internal/models"
	"cafeteria-uleam-api/internal/storage/sqlcdb"
)
type AlmacenSQLC struct {
	q *sqlcdb.Queries
}
func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{q: sqlcdb.New(db)}
}
// --- mapeo sqlc -> dominio (la capa que traduce) ---
func aEmpleadoDominio(p sqlcdb.Empleado) models.Empleado {
	return models.Empleado{
		ID:          int(p.ID),
		Nombre:      p.Nombre,
		Apellido:    p.Apellido,
		Salario:     p.Salario,
		HorasSemana: int(p.HorasSemana),
		CargoID:     int(p.CargoID),
	}
}

func aCargoDominio(p sqlcdb.Cargo) models.Cargo {
	return models.Cargo{
		ID:          int(p.ID),
		Nombre:      p.Nombre,
		Descripcion: p.Descripcion,
	}
}

func (a *AlmacenSQLC) ListarEmpleados() []models.Empleado {
	filas, err := a.q.ListarEmpleados(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.Empleado, 0, len(filas))
	for _, f := range filas {
		out = append(out, aEmpleadoDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarEmpleadoPorID(id int) (models.Empleado, bool) {
	f, err := a.q.BuscarEmpleadoPorID(context.Background(), int64(id))
	if err != nil {
		return models.Empleado{}, false // absorbe sql.ErrNoRows
	}
	return aEmpleadoDominio(f), true
}

func (a *AlmacenSQLC) CrearEmpleado(p models.Empleado) models.Empleado {
	f, err := a.q.CrearEmpleado(context.Background(), sqlcdb.CrearEmpleadoParams{
		Nombre:      p.Nombre,
		Apellido:    p.Apellido,
		Salario:     p.Salario,
		HorasSemana: int64(p.HorasSemana),
		CargoID:     int64(p.CargoID),
	})
	if err != nil {
		return models.Empleado{}
	}
	return aEmpleadoDominio(f)
}

func (a *AlmacenSQLC) ActualizarEmpleado(id int, datos models.Empleado) (models.Empleado, bool) {
	f, err := a.q.ActualizarEmpleado(context.Background(), sqlcdb.ActualizarEmpleadoParams{
		Nombre:      datos.Nombre,
		Apellido:    datos.Apellido,
		Salario:     datos.Salario,
		HorasSemana: int64(datos.HorasSemana),
		CargoID:     int64(datos.CargoID),
		ID:          int64(id),
	})
	if err != nil {
		return models.Empleado{}, false
	}
	return aEmpleadoDominio(f), true
}

func (a *AlmacenSQLC) BorrarEmpleado(id int) bool {
	filas, err := a.q.BorrarEmpleado(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

func (a *AlmacenSQLC) ListarCargos() []models.Cargo {
	filas, err := a.q.ListarCargos(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.Cargo, 0, len(filas))
	for _, f := range filas {
		out = append(out, aCargoDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarCargoPorID(id int) (models.Cargo, bool) {
	f, err := a.q.BuscarCargoPorID(context.Background(), int64(id))
	if err != nil {
		return models.Cargo{}, false
	}
	return aCargoDominio(f), true
}

func (a *AlmacenSQLC) CrearCargo(c models.Cargo) models.Cargo {
	f, err := a.q.CrearCargo(context.Background(), sqlcdb.CrearCargoParams{
		Nombre:      c.Nombre,
		Descripcion: c.Descripcion,
	})
	if err != nil {
		return models.Cargo{}
	}
	return aCargoDominio(f)
}

func (a *AlmacenSQLC) ActualizarCargo(id int, datos models.Cargo) (models.Cargo, bool) {
	f, err := a.q.ActualizarCargo(context.Background(), sqlcdb.ActualizarCargoParams{
		Nombre:      datos.Nombre,
		Descripcion: datos.Descripcion,
		ID:          int64(id),
	})
	if err != nil {
		return models.Cargo{}, false
	}
	return aCargoDominio(f), true
}

func (a *AlmacenSQLC) BorrarCargo(id int) bool {
	filas, err := a.q.BorrarCargo(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

var _ Almacen = (*AlmacenSQLC)(nil)