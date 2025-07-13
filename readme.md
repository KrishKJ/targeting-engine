# 🎯 Targeting Engine

A high-performance ad targeting engine built with Go, PostgreSQL, Redis, and Docker. It matches incoming traffic requests to active campaigns based on targeting rules like OS, app, and country.

## 🔧 Tech Stack

- Go (Golang) – Gin web framework
- PostgreSQL – Primary DB
- Redis – Smart caching layer
- Docker + Docker Compose – Containerization
- Prometheus + Grafana – Metrics and Monitoring
- pprof – Performance profiling


## 🚀 Getting Started

1. **Clone the repo**

```bash
git clone https://github.com/<your-username>/targeting-engine.git
cd targeting-engine


# Build and run containers
docker-compose up --build

# Seed database with campaigns (inside container)
docker cp migrate/seed.sql targeting-postgres:/seed.sql
# After copying file 
docker exec -it targeting-postgres psql -U postgres -d postgres -f /seed.sql


# Enter PostgreSQL container
docker exec -it targeting-postgres bash

# Run SQL
psql -U postgres -d postgres -f /seed.sql
exit

# Command to refresh cache 
curl http://localhost:8080/api/v1/refresh-cache

Command to get matched campaigns: curl "http://localhost:8080/api/v1/delivery?app=com.gametion.ludokinggame&country=us&os=android"

## 📡 API Endpoints

- `GET /api/v1/delivery?app=...&country=...&os=...` → returns matching campaigns
- `GET /api/v1/refresh-cache` → reloads data from DB into Redis
- `GET /metrics` → Prometheus metrics
- `GET /health` → Basic health check

# Sample URL
curl http://localhost:8080/api/v1/delivery?app=com.gametion.ludokinggame&country=us&os=android

# To check metrics
Prometheus: http://localhost:9090

Grafana: http://localhost:3000

Default login: admin / admin

Add Prometheus as a data source (URL: http://prometheus:9090)

Import dashboard ID: 1860 for Go metrics

# Docker Commands:
docker-compose down
docker-compose up --build

To see pprof on Browser:
http://localhost:6060/debug/pprof/