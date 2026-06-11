CREATE TABLE cargos(
     id INTEGER PRIMARY KEY,
     nombre TEXT NOT NULL,
     descripcion TEXT NOT NULL
);

CREATE TABLE empleados(
       id INTEGER PRIMARY KEY,
      nombre TEXT NOT NULL,
      apellido TEXT NOT NULL,
      salario REAL NOT NULL,
      horas_semana INTEGER NOT NULL,
      cargo_id INTEGER NOT NULL
);