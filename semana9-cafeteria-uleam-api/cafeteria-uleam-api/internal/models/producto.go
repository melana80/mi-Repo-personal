// Package models define las entidades del dominio cafetería.
package models

// Producto representa un producto a la venta en la cafetería.
//
// CategoriaID referencia el ID de una Categoria por número.
// Decisión arquitectónica: usamos ID en lugar de struct anidado.
// Esto facilita la transición a una base de datos relacional en S8 (GORM),
// donde las foreign keys son IDs, no objetos.
type Producto struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
	CategoriaID int     `json:"categoria_id"`
}
