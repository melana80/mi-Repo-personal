// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"sync"

	"github.com/joancema/examen-lavanderia/internal/models"
)

// ServicioMemoria implementa ServicioRepository en memoria.
// Se usa en los tests de reglas de negocio como fake del repositorio real.
type ServicioMemoria struct {
	mu     sync.Mutex
	datos  map[uint]models.Servicio
	nextID uint
}

func NuevoServicioMemoria() *ServicioMemoria {
	return &ServicioMemoria{datos: make(map[uint]models.Servicio), nextID: 1}
}

func (r *ServicioMemoria) Crear(h *models.Servicio) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	h.ID = r.nextID
	r.nextID++
	r.datos[h.ID] = *h
	return nil
}

func (r *ServicioMemoria) ObtenerPorID(id uint) (models.Servicio, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	h, ok := r.datos[id]
	return h, ok
}

func (r *ServicioMemoria) Listar() ([]models.Servicio, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	lista := make([]models.Servicio, 0, len(r.datos))
	for _, h := range r.datos {
		lista = append(lista, h)
	}
	return lista, nil
}

func (r *ServicioMemoria) Actualizar(h *models.Servicio) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.datos[h.ID]; !ok {
		return ErrRegistroNoExiste
	}
	r.datos[h.ID] = *h
	return nil
}
