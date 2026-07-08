package storage

import (
	"time"

	"gorm.io/gorm"

	"cafeteria-uleam-api/internal/models"

)

// Usuario representa un usuario de la cafetería.
type UsuarioGROM struct {
	db *gorm.DB
}

// NewUsuarioRepository crea un nuevo repositorio de usuarios.
func NewUsuarioRepository(db *gorm.DB) *UsuarioGROM {
	return &UsuarioGROM{db: db}
}

// CrearUsuario crea un nuevo usuario.
func (r *UsuarioGROM) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	u.CreatedAt= time.Now()
	if err := r.db.Create(&u).Error; err != nil {
		return models.Usuario{}, err
	}
	return u,nil
	
}
// BuscarUsuarioPorEmail busca un usuario por email.
func (r *UsuarioGROM) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	var u models.Usuario
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return models.Usuario{}, false
	}
	return u, true
}