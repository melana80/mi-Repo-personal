package services

import (
	"github.com/joancema/examen-lavanderia/internal/models"
	"github.com/joancema/examen-lavanderia/internal/storage"
)

type OrdenService struct {
	ordenes   storage.OrdenRepository
	servicios storage.ServicioRepository
	clientes  storage.ClienteRepository
}

func NuevaOrdenService(
	ordenes storage.OrdenRepository,
	servicios storage.ServicioRepository,
	clientes storage.ClienteRepository,
) *OrdenService {
	return &OrdenService{
		ordenes:   ordenes,
		servicios: servicios,
		clientes:  clientes,
	}
}

func (s *OrdenService) Crear(a *models.Orden) error {
	if a.Cantidad == 0 || a.ServicioID == 0 || a.ClienteID == 0 {
		return ErrDatosInvalidos
	}

	servicio, ok := s.servicios.ObtenerPorID(a.ServicioID)
	if !ok || !servicio.Activo {
		return ErrReferenciaInvalida
	}

	_, ok = s.clientes.ObtenerPorID(a.ClienteID)
	if !ok {
		return ErrReferenciaInvalida
	}
//R2: Cantidad mayor a stock disponible → ErrStockInsuficiente.
	if a.Cantidad > servicio.Stock {
		return ErrStockInsuficiente
	}
//R3: Total = precio × cantidad; descuento 10% si cantidad ≥ 5.
	precio := servicio.PrecioUnitario * float64(a.Cantidad)
	if a.Cantidad >= 5 {
		precio = precio * 0.90
	}

	servicio.Stock -= a.Cantidad
	if err := s.servicios.Actualizar(&servicio); err != nil {
		return err
	}

	a.Estado = models.EstadoPendiente
	a.Total = precio

	return s.ordenes.Crear(a)
}

func (s *OrdenService) ObtenerPorID(id uint) (models.Orden, error) {
	a, ok := s.ordenes.ObtenerPorID(id)
	if !ok {
		return models.Orden{}, ErrNoEncontrado
	}
	return a, nil
}

func (s *OrdenService) Listar() ([]models.Orden, error) {
	return s.ordenes.Listar()
}

func (s *OrdenService) Cancelar(id uint) error {
	a, ok := s.ordenes.ObtenerPorID(id)
	if !ok {
		return ErrNoEncontrado
	}
//R1: No crea orden si servicio no existe, está inactivo o cliente no existe → ErrReferenciaInvalida.
	if a.Estado != models.EstadoPendiente {
		return ErrEstadoInvalido
	}
//R4: Cancelar solo permite órdenes PENDIENTE; si están LISTA → ErrEstadoInvalido.
	servicio, ok := s.servicios.ObtenerPorID(a.ServicioID)
	if !ok {
		return ErrReferenciaInvalida
	}
//R5: Al crear orden descuenta stock; al cancelar repone exactamente la cantidad.
	servicio.Stock += a.Cantidad
	if err := s.servicios.Actualizar(&servicio); err != nil {
		return err
	}

	a.Estado = models.EstadoCancelada
	return s.ordenes.Actualizar(&a)
}
