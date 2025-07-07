# Auth Service - Microservicio de Autenticación

Este es un microservicio de autenticación desarrollado en Go que maneja el login, logout y validación de roles de usuarios.

## Características

- Autenticación de usuarios
- Gestión de sesiones con Redis
- Validación de roles
- Integración con Kafka para eventos
- API RESTful

## Estructura del Proyecto

```
.
├── cmd/
│   └── main.go
├── config/
│   └── config.go
├── internal/
│   ├── auth/
│   │   ├── handlers.go
│   │   └── middleware.go
│   ├── event/
│   │   └── kafka.go
│   └── session/
│       └── redis.go
├── tests/
│   └── auth_test.go
├── dockerfile
├── docker-compose.yml
└── .github/workflows/ci-cd.yml
```

## Configuración del CI/CD Pipeline

### Secrets Requeridos en GitHub

Para que el pipeline funcione correctamente, necesitas configurar los siguientes secrets en tu repositorio de GitHub:

1. Ve a tu repositorio en GitHub
2. Navega a **Settings** > **Secrets and variables** > **Actions**
3. Agrega los siguientes secrets:

#### `DOCKER_USERNAME`
Tu nombre de usuario de Docker Hub.

#### `DOCKER_PASSWORD`
Tu contraseña de Docker Hub o un token de acceso.

### Cómo crear un token de acceso de Docker Hub:

1. Ve a [Docker Hub](https://hub.docker.com)
2. Inicia sesión en tu cuenta
3. Ve a **Account Settings** > **Security**
4. Haz clic en **New Access Token**
5. Dale un nombre descriptivo (ej: "GitHub Actions")
6. Copia el token generado y úsalo como `DOCKER_PASSWORD`

## Pipeline de CI/CD

El pipeline incluye:

### Job de Pruebas (`test`)
- Ejecuta en Ubuntu Latest
- Configura Go 1.20
- Instala dependencias
- Ejecuta pruebas unitarias
- Ejecuta todas las pruebas del proyecto

### Job de Construcción y Despliegue (`build-and-push`)
- Se ejecuta solo después de que las pruebas pasen
- Solo se ejecuta en la rama `main`
- Construye la imagen Docker
- Sube la imagen a Docker Hub
- Actualiza la etiqueta `latest`

## Variables de Entorno

El servicio requiere las siguientes variables de entorno:

- `USER_PROFILE_SERVICE_URL`: URL del servicio de perfiles de usuario
- `REDIS_HOST`: Host de Redis
- `REDIS_PORT`: Puerto de Redis
- `REDIS_PASS`: Contraseña de Redis
- `SESSION_TTL`: TTL de las sesiones en segundos

## Ejecutar Localmente

```bash
# Instalar dependencias
go mod download

# Ejecutar pruebas
go test -v ./tests/...

# Construir y ejecutar con Docker
docker build -t auth-service .
docker run -p 8080:8080 auth-service
```

## Endpoints

- `POST /login` - Autenticación de usuario
- `POST /logout` - Cerrar sesión
- `GET /validate-role` - Validar rol de usuario

## Desarrollo

Para ejecutar las pruebas localmente:

```bash
go test -v ./tests/...
```

Para ejecutar todas las pruebas:

```bash
go test -v ./...
```
