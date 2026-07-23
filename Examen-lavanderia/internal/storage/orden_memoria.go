// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"sync"

	"github.com/joancema/examen-lavanderia/internal/models"
)

// OrdenMemoria implementa OrdenRepository en memoria.
// Se usa en los tests de reglas de negocio como fake del repositorio real.
type OrdenMemoria struct {
	mu     sync.Mutex
	datos  map[uint]models.Orden
	nextID uint
}

func NuevaOrdenMemoria() *OrdenMemoria {
	return &OrdenMemoria{datos: make(map[uint]models.Orden), nextID: 1}
}

func (r *OrdenMemoria) Crear(a *models.Orden) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	a.ID = r.nextID
	r.nextID++
	r.datos[a.ID] = *a
	return nil
}

func (r *OrdenMemoria) ObtenerPorID(id uint) (models.Orden, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.datos[id]
	return a, ok
}

func (r *OrdenMemoria) Listar() ([]models.Orden, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	lista := make([]models.Orden, 0, len(r.datos))
	for _, a := range r.datos {
		lista = append(lista, a)
	}
	return lista, nil
}

func (r *OrdenMemoria) Actualizar(a *models.Orden) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.datos[a.ID]; !ok {
		return ErrRegistroNoExiste
	}
	r.datos[a.ID] = *a
	return nil
}
