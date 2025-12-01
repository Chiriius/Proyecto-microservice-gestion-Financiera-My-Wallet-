# Proyecto-microservice-gestion-Financiera-My-Wallet-

# Historias de Usuario — Proyecto de Gestión Financiera

1. Historia de Usuario — Registro de Usuario

Como usuario nuevo
Quiero crear una cuenta en el sistema
Para acceder a mis herramientas de gestión financiera

Criterios de Aceptación:
	•	Dado que estoy en el servicio de autenticación
	•	Cuando envío mi nombre, correo y contraseña
	•	Entonces el sistema debe registrar mi usuario y generar un token JWT válido.
	•	Dado que ingreso un correo ya registrado
	•	Cuando intento crear la cuenta
	•	Entonces el sistema debe notificar que el correo ya está en uso.

⸻

2. Historia de Usuario — Inicio de Sesión

Como usuario registrado
Quiero iniciar sesión
Para acceder a mi información financiera personal

Criterios de Aceptación:
	•	Dado que mis credenciales son válidas
	•	Cuando las envío al servicio de autenticación
	•	Entonces debo recibir un token JWT activo.
	•	Dado que las credenciales son incorrectas
	•	Cuando intento iniciar sesión
	•	Entonces el sistema debe rechazar el inicio de sesión.

⸻

3. Historia de Usuario — Registrar un Gasto

Como usuario autenticado
Quiero registrar un gasto
Para llevar el control de mis finanzas diarias

Criterios de Aceptación:
	•	Dado que tengo un token JWT válido
	•	Cuando envío el monto, categoría, fecha y descripción
	•	Entonces el sistema debe guardar el gasto correctamente.
	•	Dado que falta un dato obligatorio
	•	Cuando intento registrar el gasto
	•	Entonces el sistema debe mostrar un error de validación.

⸻

4. Historia de Usuario — Registrar un Ingreso

Como usuario autenticado
Quiero registrar un ingreso
Para conocer mi balance financiero real

Criterios de Aceptación:
	•	Dado que estoy autenticado
	•	Cuando envío un monto, fecha y descripción
	•	Entonces el ingreso debe guardarse correctamente.
	•	Dado que envío un monto inválido
	•	Cuando intento registrar el ingreso
	•	Entonces el sistema debe rechazar la solicitud.

⸻

5. Historia de Usuario — Obtener Reporte Mensual

Como usuario autenticado
Quiero obtener un reporte mensual de mis ingresos y gastos
Para analizar mi salud financiera y tomar decisiones informadas

Criterios de Aceptación:
	•	Dado que estoy autenticado
	•	Cuando solicito el reporte del mes
	•	Entonces el sistema debe devolver:
	•	Total de ingresos
	•	Total de gastos
	•	Balance
	•	Categorías más usadas
	•	Dado que no existen transacciones en el mes
	•	Entonces el reporte debe mostrar valores en cero correctamente.
