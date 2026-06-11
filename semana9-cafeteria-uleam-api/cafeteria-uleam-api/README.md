# Cafetería ULEAM — Proyecto base Día B, Semana 6

API REST de la Cafetería Universitaria. Punto de partida del taller multi-entidad del Día B.

## ¿Qué hay en este proyecto?

Está implementada la entidad **Producto** completa: modelo + storage + handlers + rutas, con la misma arquitectura del Día A (Chi + estructura `cmd/`+`internal/` + struct `Server` con inyección de dependencias).

## Lo que ustedes tienen que construir

Su trabajo en este taller es agregar una **segunda entidad: `Categoria`**, siguiendo el mismo patrón que ya está implementado para Producto. La relación entre las entidades es por ID: `Producto.CategoriaID` apunta a `Categoria.ID`.

Los archivos del proyecto base tienen comentarios `// TODO TALLER:` señalando exactamente dónde deben agregar código nuevo.

## Estructura

```
cafeteria-uleam-api/
├── go.mod
├── go.sum
├── README.md
├── cmd/
│   └── cafeteria-api/
│       └── main.go              ← entry point
└── internal/
    ├── models/
    │   └── producto.go          ← Producto (ustedes agregan categoria.go)
    ├── storage/
    │   └── memoria.go           ← métodos para Producto (ustedes extienden con Categoria)
    └── handlers/
        └── producto.go          ← 5 handlers de Producto (ustedes agregan categoria.go)
```

## Requisitos

- Go 1.22 o superior
- Postman o `curl` para probar los endpoints

## Cómo correrlo

```bash
cd cafeteria-uleam-api
go mod tidy
go run ./cmd/cafeteria-api
```

Verán en consola:

```
Servidor escuchando en http://localhost:8080
```

## Productos pre-cargados al arrancar

| ID | Nombre | Precio | Stock | CategoriaID |
|----|--------|--------|-------|-------------|
| 1 | Café americano | 1.25 | 50 | 1 |
| 2 | Capuccino | 1.75 | 40 | 1 |
| 3 | Sandwich vegetariano | 2.50 | 20 | 2 |
| 4 | Croissant de jamón | 1.80 | 25 | 2 |
| 5 | Jugo de naranja | 1.50 | 30 | 3 |
| 6 | Galleta de avena | 0.75 | 60 | 4 |

**Pregunta para pensar:** ¿qué pasa si hago `GET /api/v1/productos/1` ahora? El producto referencia `CategoriaID: 1`, pero las categorías aún no existen como entidad. Esa tensión es deliberada — al final del taller, esa referencia debe resolverse en una entidad real.

## Endpoints implementados (Producto)

| Método | Ruta                          | Qué hace |
|--------|-------------------------------|----------|
| GET    | `/api/v1/productos`           | Lista todos |
| GET    | `/api/v1/productos/{id}`      | Devuelve uno |
| POST   | `/api/v1/productos`           | Crea uno |
| PUT    | `/api/v1/productos/{id}`      | Actualiza uno |
| DELETE | `/api/v1/productos/{id}`      | Elimina uno |

## Endpoints que ustedes deben implementar (Categoria)

| Método | Ruta                           | Qué hace |
|--------|--------------------------------|----------|
| GET    | `/api/v1/categorias`           | Lista todas |
| GET    | `/api/v1/categorias/{id}`      | Devuelve una |
| POST   | `/api/v1/categorias`           | Crea una |
| PUT    | `/api/v1/categorias/{id}`      | Actualiza una |
| DELETE | `/api/v1/categorias/{id}`      | Elimina una |

## Verificación con `curl`

**Listar productos:**

```bash
curl -i http://localhost:8080/api/v1/productos
```

**Obtener un producto:**

```bash
curl -i http://localhost:8080/api/v1/productos/1
```

**Crear un producto nuevo:**

```bash
curl -i -X POST http://localhost:8080/api/v1/productos \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Té verde","precio":1.00,"stock":30,"categoria_id":3}'
```

**Actualizar:**

```bash
curl -i -X PUT http://localhost:8080/api/v1/productos/1 \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Café americano grande","precio":1.50,"stock":45,"categoria_id":1}'
```

**Eliminar:**

```bash
curl -i -X DELETE http://localhost:8080/api/v1/productos/6
```

**Probar validación (nombre vacío):**

```bash
curl -i -X POST http://localhost:8080/api/v1/productos \
  -H "Content-Type: application/json" \
  -d '{"nombre":"","precio":1.00,"stock":10,"categoria_id":1}'
# Esperado: 400 Bad Request
```

**Probar validación (precio negativo):**

```bash
curl -i -X POST http://localhost:8080/api/v1/productos \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Producto raro","precio":-5,"stock":10,"categoria_id":1}'
# Esperado: 400 Bad Request
```

## Lectura recomendada antes de empezar

Antes de teclear cualquier código nuevo, **lean con calma estos tres archivos** del proyecto base. Son la plantilla a imitar:

1. `internal/models/producto.go` — cómo se declara un struct del dominio
2. `internal/storage/memoria.go` — cómo se gestiona estado en memoria con mutex
3. `internal/handlers/producto.go` — cómo se escriben handlers como métodos de Server

Si pueden explicar cada línea de esos archivos en sus propias palabras, están listos para empezar.
