# Microservicio de Usuarios (users-service-udc)

Este microservicio es uno de los cinco que componen el Sistema de Gestión de Urgencias del Café.  
Su propósito es centralizar la gestión de usuarios, incluyendo:

- **Autenticación (Auth):** Login de administradores y técnicos, emisión de tokens JWT.  
- **Gestión de Técnicos:** Creación, consulta, actualización, eliminación y cambio de contraseña.  
- **Gestión de Clientes:** administración de la información de los clientes que interactúan con el sistema.  



## Tecnologías utilizadas

- **Go 1.22+**
- **Air** (recarga automática en desarrollo)
- **Swagger** (documentación interactiva de la API)
- **Gorilla Mux** (router HTTP)


## Instalación y ejecución

### 1. Clonar el repositorio
```bash
git clone https://github.com/tu-org/user-service-ucd.git
cd user-service-ucd
```

### 2. Instalar dependencias
```bash
go mod tidy
```

### 3. Instalar Air para desarrollo
```bash
go install github.com/cosmtrek/air@latest
```

Verificar instalación:
```bash
air -v
```

### 4. Ejecutar el microservicio

Con recarga automática:
```bash
air
```

O manualmente:
```bash
go run main.go
```

El servicio quedará expuesto en:
```
http://localhost:8080
```

---

## Documentación Swagger

La API está documentada con Swagger.

Accede desde: **http://localhost:8080/swagger/index.html**

Allí encontrarás todos los endpoints organizados en tags:
- **auth** → Login y autenticación JWT
- **technicians** → CRUD y seguridad de técnicos
- **clients** → Gestión de clientes



## Endpoints disponibles

### Autenticación
- `POST /user-service/v1/auth/login` - Login de usuarios

### Gestión de Clientes
- `GET /user-service/v1/clients` - Obtener todos los clientes
- `GET /user-service/v1/clients/{clientID}` - Obtener cliente por ID
- `POST /user-service/v1/createClient` - Crear nuevo cliente
- `PUT /user-service/v1/updateClient/{clientID}` - Actualizar cliente
- `DELETE /user-service/v1/deleteClient/{clientID}` - Eliminar cliente

### Gestión de Técnicos
- `GET /user-service/v1/technicians` - Obtener todos los técnicos
- `GET /user-service/v1/technician/{technicianID}` - Obtener técnico por ID
- `POST /user-service/v1/createTechnician` - Crear nuevo técnico
- `PUT /user-service/v1/updateTechnician/{technicianID}` - Actualizar técnico
- `DELETE /user-service/v1/deleteTechnician/{technicianID}` - Eliminar técnico
- `PATCH /user-service/v1/changePassword` - Cambiar contraseña de técnico
