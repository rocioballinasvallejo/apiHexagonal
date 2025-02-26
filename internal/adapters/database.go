package adapters

import (
	"database/sql"
	"fmt"
	"monitor-pc/internal/core"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "3306"
	}
	if dbName == "" {
		dbName = "test"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error conectando a la base de datos: %v", err)
	}

	return DB.Ping()
}
func ObtenerUltimaMetrica() (*core.SystemMetrics, error) {
	query := `SELECT id, cpu_usage, ram_usage, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') 
          FROM metrics 
          ORDER BY id DESC 
          LIMIT 1`

var createdAtStr string
var metric core.SystemMetrics

err := DB.QueryRow(query).Scan(&metric.ID, &metric.CPUUsage, &metric.RAMUsage, &createdAtStr)
if err == sql.ErrNoRows {
    return nil, fmt.Errorf("no se encontraron métricas")
} else if err != nil {
    return nil, err
}

// Convertir la fecha string a time.Time
metric.Date, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
if err != nil {
    return nil, fmt.Errorf("error al convertir la fecha: %v", err)
}

return &metric, nil
}


func ObtenerMetricasDespuesDeID(lastID int) ([]core.SystemMetrics, error) {
	query := `SELECT id, cpu_usage, ram_usage, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') 
          FROM metrics 
          WHERE id > ? 
          ORDER BY id ASC`

rows, err := DB.Query(query, lastID)
if err != nil {
    return nil, err
}
defer rows.Close()

var metrics []core.SystemMetrics
for rows.Next() {
    var metric core.SystemMetrics
    var createdAtStr string

    if err := rows.Scan(&metric.ID, &metric.CPUUsage, &metric.RAMUsage, &createdAtStr); err != nil {
        return nil, err
    }

    // Convertir el string a time.Time
    metric.Date, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
    if err != nil {
        return nil, fmt.Errorf("error al convertir la fecha: %v", err)
    }

    metrics = append(metrics, metric)
}
return metrics, nil

}
