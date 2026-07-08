package storage

import "cafeteria-uleam-api/internal/models"

// Almacen define QUÉ sabe hacer un almacén de la cafetería, sin decir CÓMO.
//
// Memoria (slices) ya cumple esta interfaz sin cambios — por el duck typing
// que vimos en S3 — y AlmacenSQLite (GORM) la cumple igual. El Server depende
// de esta interfaz, no de una implementación concreta: por eso podemos cambiar
// el backend de almacenamiento sin tocar un solo handler.

//repositorio 
type ProductoRepository interface {
	ListarProductos() []models.Producto
	BuscarProductoPorID(id int) (models.Producto, bool)
	CrearProducto(p models.Producto) models.Producto
	ActualizarProducto(id int, datos models.Producto) (models.Producto, bool)
	BorrarProducto(id int) bool
}
//Repositorio de categorias
type CategoriaRepository interface {
	ListarCategorias() []models.Categoria
	BuscarCategoriaPorID(id int) (models.Categoria, bool)
	CrearCategoria(c models.Categoria) models.Categoria
	ActualizarCategoria(id int, datos models.Categoria) (models.Categoria, bool)
	BorrarCategoria(id int) bool
}
//Repositorio de usuarios
type UsuarioRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
	
}
// Almacen es un TERCER backend de la cafetería.
type Almacen interface {
	ProductoRepository
	CategoriaRepository
}

// Chequeo en tiempo de compilación: si Memoria dejara de cumplir Almacen,
// el proyecto NO compila. Red de seguridad opcional.
var _ Almacen = (*Memoria)(nil)


