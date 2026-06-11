// Command cafeteria-api arranca el servidor HTTP de la Cafetería Universitaria.
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"

	"cafeteria-uleam-api/internal/handlers"
	"cafeteria-uleam-api/internal/models"
	"cafeteria-uleam-api/internal/storage"
)

func main() {

	// 1. Abrir SQLite y migrar el esquema (crea las tablas si no existen).
	db, err := gorm.Open(sqlite.Open("cafeteria.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}
	if err := db.AutoMigrate(&models.Empleado{}, &models.Cargo{}); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	// 2. Crear el almacenamiento GORM y sembrar si está vacío.
	//    (es decir, crear los cargos y empleados iniciales).
	almacenGorm := storage.NuevoAlmacenSQLite(db)
	almacenGorm.SembrarSiVacio()

	var almacen storage.Almacen
	switch os.Getenv("STORAGE") {
	case "sqlc":
		log.Fatal("sqlc: ", err)
		if sqlDB, err := db.DB(); err == nil {
			_ = sqlDB.Close() // GORM ya migró y sembró; liberamos su conexión
		}
		sdb, err := sql.Open("sqlite", "cafeteria.db")
		if err != nil {
			log.Fatal("sqlc: ", err)
		}
		almacen = storage.NuevoAlmacenSQLC(sdb)
	default:
		almacen = almacenGorm
	}

	// 3. Server con inyección de dependencias (recibe un storage.Almacen).
	servidor := handlers.NewServer(almacen) // idéntico: no sabe cuál recibió

	// 4. Router + middleware.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: false,
	}))

	// 5. Rutas versionadas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/empleados", servidor.ListarEmpleados)
		r.Post("/empleados", servidor.CrearEmpleado)
		r.Get("/empleados/{id}", servidor.ObtenerEmpleado)
		r.Put("/empleados/{id}", servidor.ActualizarEmpleado)
		r.Delete("/empleados/{id}", servidor.BorrarEmpleado)

		r.Get("/cargos", servidor.ListarCargos)
		r.Post("/cargos", servidor.CrearCargo)
		r.Get("/cargos/{id}", servidor.ObtenerCargo)
		r.Put("/cargos/{id}", servidor.ActualizarCargo)
		r.Delete("/cargos/{id}", servidor.BorrarCargo)
	})

	log.Println("Servidor escuchando en http://localhost:8080")

	// Servir frontend estático
	frontendPath := "frontend"
	if _, err := os.Stat(frontendPath); err == nil {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, filepath.Join(frontendPath, "index.html")) //conectamos el servidor con el frontend
		})
	}
	log.Fatal(http.ListenAndServe(":8080", r))
}
