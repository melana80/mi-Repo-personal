-- name: ListarEmpleados :many
SELECT id, nombre, apellido, salario, horas_semana, cargo_id FROM empleados;
-- name: BuscarEmpleadoPorID :one
SELECT id, nombre, apellido, salario, horas_semana, cargo_id FROM empleados
WHERE id = ?;
-- name: CrearEmpleado :one
INSERT INTO empleados (nombre, apellido, salario, horas_semana, cargo_id)
VALUES (?, ?, ?, ?, ?)
RETURNING id, nombre, apellido, salario, horas_semana, cargo_id;
-- name: ActualizarEmpleado :one
UPDATE empleados
SET nombre = ?, apellido = ?, salario = ?, horas_semana = ?, cargo_id = ?
WHERE id = ?
RETURNING id, nombre, apellido, salario, horas_semana, cargo_id;
-- name: BorrarEmpleado :execrows
DELETE FROM empleados WHERE id = ?;

-- name: ListarCargos :many
SELECT id, nombre, descripcion FROM cargos;
-- name: BuscarCargoPorID :one
SELECT id, nombre, descripcion FROM cargos
WHERE id = ?;
-- name: CrearCargo :one
INSERT INTO cargos (nombre, descripcion)
VALUES (?, ?)
RETURNING id, nombre, descripcion;
-- name: ActualizarCargo :one
UPDATE cargos
SET nombre = ?, descripcion = ?
WHERE id = ?
RETURNING id, nombre, descripcion;
-- name: BorrarCargo :execrows
DELETE FROM cargos WHERE id = ?;