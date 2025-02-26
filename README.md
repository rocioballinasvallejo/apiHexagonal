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
