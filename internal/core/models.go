package core

import "time"

// SystemMetric representa una medición de recursos del sistema
type SystemMetric struct {
	ID       int       `json:"id"`
	CPUUsage float64   `json:"cpu_usage"`
	RAMUsage float64   `json:"ram_usage"`
	Date     time.Time `json:"created_at"`

}

// SystemMetrics representa las métricas obtenidas en tiempo real
type SystemMetrics struct {
	ID       int       `json:"id"`
	CPUUsage float64   `json:"cpu_usage"`
	RAMUsage float64   `json:"ram_usage"`
	Date     time.Time `json:"created_at"`

}

// MetricsDB representa cómo se almacena en MySQL
type MetricsDB struct {
	ID        int       `json:"id"`
	CPUUsage  float64   `json:"cpu_usage"`
	RAMUsage  float64   `json:"ram_usage"`
	CreatedAt time.Time `json:"created_at"`
}

// AlertDB representa las alertas en MySQL
type AlertDB struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	AlertType string    `json:"alert_type"`
	CPUUsage  float64   `json:"cpu_usage"`
	RAMUsage  float64   `json:"ram_usage"`
	CreatedAt time.Time `json:"created_at"`
	Date      time.Time `json:"date"`
}
