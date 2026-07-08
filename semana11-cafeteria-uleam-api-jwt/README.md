# SOLUCIÓN DE REFERENCIA — Taller sqlc (solo docente)

API de la Cafetería con **tres backends** detrás de la misma interfaz
`storage.Almacen`: `Memoria`, `AlmacenSQLite` (GORM) y **`AlmacenSQLC`** (nuevo).
Incluye el frontend de referencia y el CORS ya re-agregado.

## Estructura agregada respecto al proyecto base

```
sqlc.yaml                          # config de sqlc
db/schema.sql                      # esquema (solo para codegen; NO se ejecuta)
db/queries.sql                     # 10 queries anotadas
internal/storage/sqlcdb/           # paquete GENERADO por sqlc
internal/storage/sqlc.go           # AlmacenSQLC (adaptador) — escrito a mano
internal/middleware/cors.go        # CORS re-agregado
frontend/index.html                # CRUD de referencia (vanilla JS)
```

> ⚠️ El paquete `internal/storage/sqlcdb/` está escrito **a mano** para igualar
> lo que produce `sqlc generate` (el entorno de generación no pudo instalar el
> binario sqlc). Antes de clase, corré `sqlc generate` vos mismo y reemplazá esa
> carpeta por la salida real — deberían ser casi idénticas.

## Cómo correr

```bash
go mod tidy
go build ./... && go vet ./...

# Backend GORM (por defecto)
go run ./cmd/cafeteria-api

# Backend sqlc (misma API, mismo frontend)
STORAGE=sqlc go run ./cmd/cafeteria-api              # macOS / Linux
# Windows PowerShell:  $env:STORAGE="sqlc"; go run ./cmd/cafeteria-api
# Windows cmd:         set STORAGE=sqlc && go run ./cmd/cafeteria-api
```

Frontend (en OTRO origen, para que el navegador aplique CORS):

```bash
cd frontend && python3 -m http.server 5500
# abrir http://localhost:5500
```

## ⚠️ Validación local CRÍTICA: el driver database/sql

El backend sqlc hace `sql.Open("sqlite", "cafeteria.db")`. Ese nombre de driver
("sqlite") lo registra `github.com/glebarez/go-sqlite`, que ya entra al grafo
porque `glebarez/sqlite` (el driver de GORM) lo usa por debajo. Por eso lo
importamos en blanco en `main.go` y **no hace falta `modernc.org/sqlite`**
(importar modernc además causaría un pánico de "Register called twice for sqlite").

Verificá en tu máquina:
- `go mod tidy` resuelve `github.com/glebarez/go-sqlite` (ajustá la versión si
  hace falta; tidy la corrige).
- Si `sql.Open("sqlite", ...)` diera "unknown driver", confirmá el nombre que
  registra tu versión de `glebarez/go-sqlite`.

## Notas pedagógicas (resumen; el detalle está en la guía docente)

- El adaptador resuelve 3 fricciones de sqlc: inyecta `context`, mapea tipos
  (`int64`↔`int`) y absorbe `error`→`bool`.
- `CrearProducto`/`CrearCategoria` no pueden reportar fallo (la interfaz no lo
  permite). Es una debilidad **heredada** de la interfaz, no introducida por sqlc.
- Quitar CORS no rompe `curl`/Postman (CORS lo aplica el navegador). Por eso el
  fallo solo aparece con el frontend. El log del servidor muestra el preflight
  fallando: `OPTIONS /api/v1/productos ... 405`.
