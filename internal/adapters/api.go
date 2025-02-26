package adapters

import (
	"strconv"
	"time"
	"fmt"
	"monitor-pc/internal/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// SetupRoutes configura las rutas del servidor
func SetupRoutes(app *fiber.App) {
	// Configurar CORS para permitir todas las solicitudes
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/metrics", GetMetrics)
	app.Post("/metrics", SaveMetric)
	app.Delete("/metrics/:id", DeleteMetric)

	app.Get("/alerts", GetAlerts)
	app.Post("/alerts", SaveAlert)
	app.Delete("/alerts/:id", DeleteAlert)

	// Nuevas rutas para polling
	app.Get("/metrics/last", GetLastMetric)   // Short polling
	app.Get("/metrics/stream", StreamMetrics) // Long polling
}

// Obtener métricas desde la base de datos
func GetMetrics(c *fiber.Ctx) error {
	metrics, err := ObtenerMetricas()
	if (err != nil) {
		return c.Status(500).JSON(fiber.Map{"error": "Error obteniendo métricas"})
	}
	return c.JSON(metrics)
}

// Guardar una métrica en la base de datos
func SaveMetric(c *fiber.Ctx) error {
	var metric core.SystemMetrics
	if err := c.BodyParser(&metric); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	err := GuardarMetrica(metric)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error guardando métrica"})
	}

	return c.JSON(fiber.Map{"message": "Métrica guardada"})
}

// Eliminar una métrica
func DeleteMetric(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	err = EliminarMetrica(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error eliminando métrica"})
	}

	return c.JSON(fiber.Map{"message": "Métrica eliminada"})
}

// Obtener alertas desde la base de datos
func GetAlerts(c *fiber.Ctx) error {
	alerts, err := ObtenerAlertas()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error obteniendo alertas"})
	}
	return c.JSON(alerts)
}

// Guardar una alerta en la base de datos
func SaveAlert(c *fiber.Ctx) error {
	var alert core.AlertDB
	if err := c.BodyParser(&alert); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	err := GuardarAlerta(alert)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error guardando alerta"})
	}

	return c.JSON(fiber.Map{"message": "Alerta guardada"})
}

// Eliminar una alerta
func DeleteAlert(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	err = EliminarAlerta(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error eliminando alerta"})
	}

	return c.JSON(fiber.Map{"message": "Alerta eliminada"})
}

// GetLastMetric implementa short polling
func GetLastMetric(c *fiber.Ctx) error {
	metric, err := ObtenerUltimaMetrica()
	if err != nil {
		fmt.Println("Error obteniendo última métrica:", err)
	} else {
		fmt.Println("Última métrica obtenida:", metric)
	}
	return c.JSON(metric)
}

// StreamMetrics implementa long polling
func StreamMetrics(c *fiber.Ctx) error {
	// Obtener el último ID conocido del cliente
	lastID, err := strconv.Atoi(c.Query("last_id", "0"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	// Esperar hasta 30 segundos por nuevas métricas
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			return c.JSON(fiber.Map{"message": "No hay nuevas métricas"})
		case <-ticker.C:
			metrics, err := ObtenerMetricasDespuesDeID(lastID)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Error obteniendo métricas"})
			}
			if len(metrics) > 0 {
				return c.JSON(metrics)
			}
		}
	}
}
