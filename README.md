# cantera-rosa

## **Overview:**

Un negocio que se dedica a rentar juegos de mesa y acompañarlos con bebidas y botanas desea llevar la gestion de su negocio.

- Requiere un portal para administras su recursos y trabajadores.
- En el portal del negocio podra gestionar sus recursos y trabajadores
- Los trabajadores podran gestionar el negocio con menos privilegios

### **Scope**

Un portal para administrar los alimentos y juegos de mesa un **CRUD** (crear, consultar, actualizar y eliminar). Tambien se podran dar de alta empleados.

Un portal para consultar y hacer pedidos.

### **Casos de uso**

La acciones permitidas a los perfiles son.

- Como trabajador y administrador quiero poder accerder con mi usuario y contraseña
- Como administrador quiero crear/editar/consultar/eliminar sucursales de mi negocio.
- Como administrador quiero crear/editar/consultar/eliminar productos de stock
- Como administrador quiero ver una bitacora/historial de los movimientos en stock
- Como administrador quiero crear/editar/consultar/eliminar productos del menu
- Como administrador quiero ver una bitacora/historial de ventas
- Como administrador quiero registrar gastos
- Como trabajador quiero ver productos
- Como trabajador quiero ver alimentos
- Como trabajador quiero crear cuentas de mesas activas
- Como trabajador quiero agregar mas productos a las cuentas activas 
- Como trabajador quiero consultar la informacion de las mesas activas
- Como trabajador quiero cerrar las cuentas activas
- Como cajero quiero hacer cortes de caja
- El sistema debe generar un historial de lo que se ha pedido

### **Out of Scope**

En esta primera version las cosas que estan fuera, pero seran implementadas en un futuro son.

- Como administrador quiero ver las estadisticas de venta
- Como administrador quiero generar encuestas 
- Como cliente quiero enviar feedback a los trabajadores
- Como cliente quiero agregar comentarios a mis pedidos

- Los clientes podran consultar el menu de alimentos y la lista de juegos disponibles
- Los clientes podran hacer pedidos por mesa y consultar su cuenta.
- asi como un espacio para sus clientes.
- Como administrador quiero generar codigos QR unicos por mesa
- Como cliente quiero consultar el menu sin estar registrado
- Como cliente quiero hacer pedidos sin estar registrado
- Como trabajador quiero recibir los pedidos con la informacion necesaria (Productos y mesa)
- Como cliente quiero poder consultar mi cuenta
- Como cliente quiero poder agregar mas productos a mi mesa
- Como cliente quiero filtrar y ordenar los juegos de mesa por: tiempo promedio de partida, numero de jugadores, edad recomendada y precio
- Como trabajador quiero recibir una notificacion de que un pedido se ah hecho.
- Como trabajador quier cerrar cuentas

---

## **Arquitectura**

### **Diagramas**

![Infraestructure](./img/cantera_infraestructura.jpeg?raw=true)

### **Modelo de datos**

Los modelos estaran en los proyectos a los que pertenecen
