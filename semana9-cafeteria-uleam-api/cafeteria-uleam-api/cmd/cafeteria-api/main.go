// Command cafeteria-api arranca el servidor HTTP de la Cafetería Universitaria.
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware" 

	"cafeteria-uleam-api/internal/handlers"
	"cafeteria-uleam-api/internal/storage"
	"cafeteria-uleam-api/internal/midleware"
)

func main() {
	// 1. Crear el almacenamiento y cargar datos iniciales.
	almacen := storage.NuevaMemoria()
	almacen.SeedProductos()
	almacen.SeedCategorias()

	// 2. Crear el Server inyectándole el almacenamiento.
	servidor := handlers.NewServer(almacen)

	// 3. Configurar el router con versionado /api/v1/.
	r := chi.NewRouter()
	r.Use(chimw.Logger) //es para ver las peticiones
	r.Use(chimw.Recoverer) //es para manejar errores
	r.Use(midleware.CORS) //es para permitir peticiones de cualquier origen

	r.Route("/api/v1", func(r chi.Router) {
		// Productos: CRUD completo.
		r.Get("/productos", servidor.ListarProductos)
		r.Post("/productos", servidor.CrearProducto)
		r.Get("/productos/{id}", servidor.ObtenerProducto)
		r.Put("/productos/{id}", servidor.ActualizarProducto)
		r.Delete("/productos/{id}", servidor.BorrarProducto)

		// Categorías: CRUD completo.
		r.Get("/categorias", servidor.ListarCategorias)
		r.Post("/categorias", servidor.CrearCategoria)
		r.Get("/categorias/{id}", servidor.ObtenerCategoria)
		r.Put("/categorias/{id}", servidor.ActualizarCategoria)
		r.Delete("/categorias/{id}", servidor.BorrarCategoria)

		// Otros endpoints.
		r.Get("/provocarerrores", func(w http.ResponseWriter, _ *http.Request) {
			panic("Error probocado desde el sevidor")
		
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080") //nolint
	log.Fatal(http.ListenAndServe(":8080", r))
}
