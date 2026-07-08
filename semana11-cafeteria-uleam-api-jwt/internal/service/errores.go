package service

import "errors"

var (
	ErrNombreVacio          = errors.New("el campo nombre es obligatorio")
	ErrNoEncontrado =        errors.New("registro no encontrado")
	ErrPrecioNegativo       = errors.New("el precio no puede ser negativo")
	ErrEmailEnUso           = errors.New("el correo electrónica ya se encuentra en uso")
	ErrCredencialesInvalidas = errors.New("Email o contraseña incorrectos")
)