////////////esquema mysql
create database test; 
use test;
CREATE TABLE metrics (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cpu_usage FLOAT NOT NULL,
    ram_usage FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE alerts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    message VARCHAR(255) NOT NULL,
    cpu_usage FLOAT NOT NULL,
    ram_usage FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
/////////////////////////////////////////////////////////////////
Probar la ruta GET /metrics
curl -X GET http://localhost:3000/metrics


Probar la ruta POST /metrics
curl -X POST http://localhost:3000/metrics -H "Content-Type: application/json" -d '{"cpu_usage": 20, "memory_usage": 30}'


Probar la ruta DELETE /metrics/:id
curl -X DELETE http://localhost:3000/metrics/1

Probar la ruta GET /alerts
curl -X GET http://localhost:3000/alerts

Probar la ruta POST /alerts
curl -X POST http://localhost:3000/alerts -H "Content-Type: application/json" -d '{"type": "CPU", "threshold": 80}'


Probar la ruta DELETE /alerts/:id
curl -X DELETE http://localhost:3000/alerts/1


Probar la ruta GET /metrics/last
curl -X GET http://localhost:3000/metrics/last


Probar la ruta GET /metrics/stream
curl -X GET http://localhost:3000/metrics/stream?last_id=0

///////////// instalar dependencias ///////

go get github.com/gofiber/fiber/v2
go get github.com/go-sql-driver/mysql
go get github.com/gofiber/fiber/v2/middleware/cors
