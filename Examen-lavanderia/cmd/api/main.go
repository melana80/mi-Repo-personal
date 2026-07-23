// ARCHIVO BLOQUEADO — NO MODIFICAR
package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/joancema/examen-lavanderia/internal/handlers"
	"github.com/joancema/examen-lavanderia/internal/models"
	"github.com/joancema/examen-lavanderia/internal/services"
	"github.com/joancema/examen-lavanderia/internal/storage"
)

func main() {
	db, err := gorm.Open(sqlite.Open("lavanderia.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("no se pudo abrir la base de datos: %v", err)
	}

	if err := db.AutoMigrate(
		&models.Servicio{},
		&models.Cliente{},
		&models.Orden{},
	); err != nil {
		log.Fatalf("error en la migración: %v", err)
	}

	sembrarServicios(db)

	// Repositories (GORM)
	servicioRepo := storage.NuevoServicioGORM(db)
	clienteRepo := storage.NuevoClienteGORM(db)
	ordenRepo := storage.NuevaOrdenGORM(db)

	// Services
	servicioSvc := services.NuevoServicioService(servicioRepo)
	clienteSvc := services.NuevoClienteService(clienteRepo)
	ordenSvc := services.NuevaOrdenService(ordenRepo, servicioRepo, clienteRepo)

	// Handlers + Router
	router := handlers.NuevoRouter(
		handlers.NuevoServicioHandler(servicioSvc),
		handlers.NuevoClienteHandler(clienteSvc),
		handlers.NuevaOrdenHandler(ordenSvc),
	)

	log.Println("API de la lavandería escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

// sembrarServicios carga el catálogo inicial solo si la tabla está vacía.
// Los clientes y ordenes se crean vía API.
func sembrarServicios(db *gorm.DB) {
	var total int64
	db.Model(&models.Servicio{}).Count(&total)
	if total > 0 {
		return
	}
	iniciales := []models.Servicio{
		{Nombre: "Lavado básico", PrecioUnitario: 8.50, Stock: 10, Activo: true},
		{Nombre: "Lavado en seco", PrecioUnitario: 6.00, Stock: 4, Activo: true},
		{Nombre: "Planchado express", PrecioUnitario: 5.00, Stock: 2, Activo: true},
		{Nombre: "Lavado de alfombras", PrecioUnitario: 15.00, Stock: 3, Activo: false},
	}
	for i := range iniciales {
		db.Create(&iniciales[i])
	}
}
