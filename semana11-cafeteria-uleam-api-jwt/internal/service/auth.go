package service

import (
	
	"strings"
	"time"

	"cafeteria-uleam-api/internal/models"
	"cafeteria-uleam-api/internal/storage"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Secret JWT
var secretJwt = []byte("essecreto")

// Duración del JWT
var duracionJwt = time.Hour * 24

// Claims es el payload del JWT de un usuario.
type Claims struct {
	UsuarioID int `json:"id"`
	jwt.RegisteredClaims
}
type AutenticacionService struct {
	repo storage.UsuarioRepository

}

func NuevaAutenticacionService(repo storage.UsuarioRepository) *AutenticacionService {
	return &AutenticacionService{repo: repo}
}

// Registrar un nuevo usuario
func (s *AutenticacionService) Registrar(email, password string) (models.Usuario, error){
	email = strings.TrimSpace(strings.ToUpper(email))
	if email == "" || strings.TrimSpace(password) ==""{
		return models.Usuario{}, ErrNombreVacio
	}
	if _, existe := s.repo.BuscarUsuarioPorEmail(email); existe{
		return models.Usuario{}, ErrEmailEnUso
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.Usuario{}, err
	}
	return s. repo.CrearUsuario(
		models.Usuario{
			Email: email,
			PasswordHash: string(hash),
			CreatedAt: time.Now(),
		},

	)
}

// Login de un usuario
func (s *AutenticacionService) Loginn(email string, password string) (string, error) {
    email = strings.TrimSpace(strings.ToUpper(email))
    u, existe := s.repo.BuscarUsuarioPorEmail(email)
	if !existe {
        return "", ErrEmailEnUso
    }
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", ErrCredencialesInvalidas
	}
    return s.generarToken(u)
}

// Generar token
func (s *AutenticacionService) generarToken(u models.Usuario) (string, error) {
	claims := Claims{
		UsuarioID: u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duracionJwt)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretJwt)
}

// Validar token
func (s *AutenticacionService) ValidarToken(tokenStr string) (int, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, ErrCredencialesInvalidas
        }
        return secretJwt, nil
    })
    if err != nil || !token.Valid {
        return 0, ErrCredencialesInvalidas
    }
    claims, ok := token.Claims.(*Claims)
    if !ok {
        return 0, ErrCredencialesInvalidas
    }
    return claims.UsuarioID, nil
}