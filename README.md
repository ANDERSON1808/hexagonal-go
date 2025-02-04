# 🚀 **Hexagonal Golang API**

API en **Golang** con arquitectura **Hexagonal (Ports & Adapters)**, basada en **DDD, SOLID, Concurrency** y **GORM** como ORM para PostgreSQL.

---

## 📌 **Índice**
- [📦 Estructura del Proyecto](#-estructura-del-proyecto)
- [⚡ Requisitos Previos](#-requisitos-previos)
- [🚀 Cómo Ejecutarlo Localmente](#-cómo-ejecutarlo-localmente)
- [🐳 Ejecutar con Docker](#-ejecutar-con-docker)
- [🛠 Endpoints de la API](#-endpoints-de-la-api)
- [📝 Documentación con Swagger](#-documentación-con-swagger)
- [✅ Ejecutar Pruebas](#-ejecutar-pruebas)

---

## 📦 **Estructura del Proyecto**

```
hexagonal-go/
│── cmd/                   # Punto de entrada de la aplicación
│── internal/
│   ├── application/       # Casos de uso (Servicios)
│   ├── domain/            # Modelos de dominio y repositorios
│   ├── infrastructure/
│   │   ├── db/            # Conexión y migraciones de la BD
│   │   ├── http/          # Controladores HTTP
│   │   ├── routes/        # Definición de rutas
│   ├── mocks/             # Mock para pruebas
│── migrations/            # Archivos de migración de BD
│── .env                   # Variables de entorno
│── docker-compose.yml     # Configuración de Docker
│── Dockerfile             # Imagen de Docker para la API
│── go.mod                 # Dependencias del proyecto
│── README.md              # Documentación
```

---

## ⚡ **Requisitos Previos**
Antes de comenzar, asegúrate de tener instalados:

✅ **[Go 1.22+](https://go.dev/dl/)**  
✅ **[Docker](https://www.docker.com/)**  
✅ **[Docker Compose](https://docs.docker.com/compose/install/)**  
✅ **[PostgreSQL](https://www.postgresql.org/)** *(Si deseas correrlo sin Docker)*

---

## 🚀 **Cómo Ejecutarlo Localmente**
1️⃣ **Clonar el repositorio**
```sh
git clone https://github.com/ANDERSON1808/hexagonal-go.git
cd hexagonal-go
```

2️⃣ **Crear un archivo `.env` con las variables de entorno**
```sh
cp .env.example .env
```
✍ **Editar `.env` con la configuración de PostgreSQL**:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=hexagonal_db
```

3️⃣ **Ejecutar las migraciones**
```sh
go run cmd/migrate.go
```

4️⃣ **Iniciar el servidor**
```sh
go run cmd/main.go
```
📌 **La API estará disponible en:** `http://localhost:8080`

---

## 🐳 **Ejecutar con Docker**
1️⃣ **Construir la imagen de la API**
```sh
docker build -t hexagonal-go .
```

2️⃣ **Levantar los contenedores con `docker-compose`**
```sh
docker-compose up -d
```

📌 **Esto ejecutará**:  
✅ API en `http://localhost:8080`  
✅ PostgreSQL en `localhost:5432`

Para detener los contenedores:
```sh
docker-compose down
```

---

## 🛠 **Endpoints de la API**
📌 **Usuarios**
| Método | Endpoint             | Descripción |
|--------|----------------------|-------------|
| `POST`   | `/users`            | Crear usuario |
| `GET`    | `/users/all`        | Obtener todos los usuarios |
| `GET`    | `/users/{id}`       | Obtener un usuario por ID |
| `DELETE` | `/users/{id}`       | Eliminar usuario |
| `POST`   | `/users/concurrent` | Crear usuarios concurrentemente |

---

## 📝 **Documentación con Swagger**
📌 La API tiene **documentación Swagger** generada automáticamente.  
Accede desde:  
👉 `http://localhost:8080/swagger/index.html`

---

## ✅ **Ejecutar Pruebas**
Ejecuta los **tests unitarios** con:
```sh
go test ./internal/... -v
```

📌 **Para ver cobertura de código**:
```sh
go test ./internal/... -cover
```

---

## 🎯 **Conclusión**
✔️ API robusta con **arquitectura hexagonal** y **DDD**.  
✔️ **GORM** para persistencia en **PostgreSQL**.  
✔️ Soporte para **concurrencia** y **migraciones**.  
✔️ **Docker & Docker Compose** para fácil despliegue.  
✔️ **Swagger** para documentación de la API.

🚀 **¡Listo para producción!** 🏆