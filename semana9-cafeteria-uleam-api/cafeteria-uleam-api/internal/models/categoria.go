package models

// Categoria representa una categoría de productos en la cafetería.
//
// La relación con Producto es por ID: Producto.CategoriaID apunta a Categoria.ID.
type Categoria struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}
