package main

import (
	"log"
	"monitor-pc/internal/adapters"
	"monitor-pc/internal/core"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func collectMetrics(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		metrics, err := core.GetSystemMetrics()
		if err != nil {
			log.Printf("Error obteniendo métricas: %v", err)
			continue
		}

		err = adapters.GuardarMetrica(*metrics)
		if err != nil {
			log.Printf("Error guardando métrica: %v", err)
		} else {
			log.Println("✅ Métrica guardada correctamente en MySQL")
		}
	}
}

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: No se pudo cargar el archivo .env: %v", err)
	}

	// Inicializar la base de datos
	if err := adapters.InitDB(); err != nil {
		log.Fatalf("Error inicializando la base de datos: %v", err)
	}

	// Iniciar recolección periódica de métricas (cada 1 minuto)
	go collectMetrics(1 * time.Minute)

	// Iniciar el servidor
	app := fiber.New()
	adapters.SetupRoutes(app)

	log.Println("Servidor ejecutándose en http://localhost:3000")
	app.Listen(":3000")
}
