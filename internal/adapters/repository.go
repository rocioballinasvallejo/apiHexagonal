package adapters

import (
	_"database/sql"
	"fmt"
	"time"
	"monitor-pc/internal/core"
)

// GuardarMetrica almacena una métrica en la base de datos
func GuardarMetrica(m core.SystemMetrics) error {
	query := "INSERT INTO metrics (cpu_usage, ram_usage) VALUES (?, ?)"
	_, err := DB.Exec(query, m.CPUUsage, m.RAMUsage)
	if err != nil {
		return fmt.Errorf("error insertando métrica: %v", err)
	}
	return nil
}


// ObtenerMetricas obtiene todas las métricas de la base de datos
func ObtenerMetricas() ([]core.MetricsDB, error) {
    rows, err := DB.Query("SELECT id, cpu_usage, ram_usage, created_at FROM metrics")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var metrics []core.MetricsDB
    for rows.Next() {
        var metric core.MetricsDB
        var createdAt string
        if err := rows.Scan(&metric.ID, &metric.CPUUsage, &metric.RAMUsage, &createdAt); err != nil {
            return nil, err
        }
        metric.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
        if err != nil {
            return nil, err
        }
        metrics = append(metrics, metric)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return metrics, nil
}

// EliminarMetrica borra una métrica por ID
func EliminarMetrica(id int) error {
	_, err := DB.Exec("DELETE FROM metrics WHERE id = ?", id)
	return err
}

// GuardarAlerta almacena una alerta en la base de datos
func GuardarAlerta(alert core.AlertDB) error {
	query := "INSERT INTO alerts (message, alert_type) VALUES (?, ?)"
	_, err := DB.Exec(query, alert.Message, alert.AlertType)
	return err
}

// ObtenerAlertas obtiene todas las alertas registradas
func ObtenerAlertas() ([]core.AlertDB, error) {
	rows, err := DB.Query("SELECT id, message, alert_type, created_at FROM alerts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []core.AlertDB
	for rows.Next() {
		var a core.AlertDB
		err := rows.Scan(&a.ID, &a.Message, &a.AlertType, &a.Date)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, a)
	}
	return alerts, nil
}

// EliminarAlerta borra una alerta por ID
func EliminarAlerta(id int) error {
	_, err := DB.Exec("DELETE FROM alerts WHERE id = ?", id)
	return err
}



// Obtener la última métrica desde la base de datos
// func ObtenerUltimaMetrica() (core.MetricsDB, error) {
//     var metric core.MetricsDB
//     row := db.QueryRow("SELECT id, cpu_usage, ram_usage, created_at FROM metrics ORDER BY created_at DESC LIMIT 1")
//     var createdAt string
//     if err := row.Scan(&metric.ID, &metric.CPUUsage, &metric.RAMUsage, &createdAt); err != nil {
//         return metric, err
//     }
//     metric.CreatedAt, err := time.Parse("2006-01-02 15:04:05", createdAt)
//     if err != nil {
//         return metric, err
//     }
//     return metric, nil
// }