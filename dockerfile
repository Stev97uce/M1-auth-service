# Etapa de construcción
FROM golang:1.20-alpine AS builder

# Instalar dependencias del sistema
RUN apk add --no-cache git

WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar código fuente
COPY . .

# Construir la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth-service ./cmd

# Etapa de ejecución
FROM alpine:latest

# Instalar ca-certificates para HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar el binario desde la etapa de construcción
COPY --from=builder /app/auth-service .

# Exponer puerto
EXPOSE 8080

# Comando por defecto
CMD ["./auth-service"]
