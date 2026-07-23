# DECISIONES

1. Servicio = catálogo de la lavandería (pantalla 01). Tiene stock y precio.
2. Cliente se relaciona con Orden (FK cliente_id) porque cada orden pertenece a un cliente (pantalla 03).
3. Orden referencia a Servicio y Cliente (pantallas 02 y 03): al crearse descuenta stock (R5) y aplica descuento 10% si cantidad ≥ 5 (R3).
