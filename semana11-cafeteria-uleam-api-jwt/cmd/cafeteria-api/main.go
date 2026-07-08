// Command cafeteria-api arranca el servidor HTTP de la Cafetería Universitaria.
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite" // driver database/sql "sqlite" (pure-Go) para el backend sqlc
	"github.com/glebarez/sqlite"      // driver GORM (pure-Go)
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"cafeteria-uleam-api/internal/handlers"
	"cafeteria-uleam-api/internal/middleware"
	"cafeteria-uleam-api/internal/models"
	"cafeteria-uleam-api/internal/storage"

	"cafeteria-uleam-api/internal/service"
)

func main() {
	// 1. GORM es el DUEÑO DEL ESQUEMA: abre la DB, migra y siembra.
	//    Esto corre siempre, sin importar qué backend sirva después.
	gdb, err := gorm.Open(sqlite.Open("cafeteria.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}
	if err := gdb.AutoMigrate(&models.Producto{}, &models.Categoria{}, &models.Usuario{}); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}
	almacenGorm := storage.NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	// 2. Elegir el backend que SIRVE las peticiones según la variable STORAGE.
	//    >>> Esta es la ÚNICA decisión que cambia entre GORM y sqlc. <<<
	var almacen storage.Almacen
	switch os.Getenv("STORAGE") {
	case "sqlc":
		// Ya migramos y sembramos con GORM; cerramos esa conexión para que
		// sqlc sea el único dueño del archivo cafeteria.db en tiempo de servicio.
		if sqlDB, err := gdb.DB(); err == nil {
			_ = sqlDB.Close()
		}
		sdb, err := sql.Open("sqlite", "cafeteria.db")
		if err != nil {
			log.Fatal("no se pudo abrir sql.DB para sqlc: ", err)
		}
		almacen = storage.NuevoAlmacenSQLC(sdb)
		log.Println("Backend de almacenamiento: sqlc (database/sql)")
	default:
		almacen = almacenGorm
		log.Println("Backend de almacenamiento: GORM")
	}

	// 3. Server con inyección de dependencias. No sabe qué backend recibió.
	usuarioRepo := storage.NewUsuarioRepository(gdb)
	authService := service.NuevaAutenticacionService(usuarioRepo)
	productoService := service.NuevaProductoService(almacen)
	categoriaService := service.NuevaCategoriaService(almacen)
	servidor := handlers.NewServer(productoService, categoriaService, authService)

	// 4. Router + middleware.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 5. Rutas versionadas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {
		// 6. Rutas sin autenticación.
		r.Post("/auth/registrar", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		r.Group(func(r chi.Router) {
			r.Get("/productos", servidor.ListarProductos)
			r.Post("/productos", servidor.CrearProducto)
			r.Get("/productos/{id}", servidor.ObtenerProducto)
			r.Put("/productos/{id}", servidor.ActualizarProducto)
			r.Delete("/productos/{id}", servidor.BorrarProducto)
		

			r.Get("/categorias", servidor.ListarCategorias)
			r.Post("/categorias", servidor.CrearCategoria)
			r.Get("/categorias/{id}", servidor.ObtenerCategoria)
			r.Put("/categorias/{id}", servidor.ActualizarCategoria)
			r.Delete("/categorias/{id}", servidor.BorrarCategoria)

		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}