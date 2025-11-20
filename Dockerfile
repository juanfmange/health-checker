# -----------------------------
# Etapa 1: Build
# -----------------------------
FROM golang:1.22-alpine AS builder

# Intalar dependencias requeridas para el build
RUN apk add --no-cache git

WORKDIR /app

# copiar el go.mod y go.sum para las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el codigo
COPY . .

# Construir el binario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o health-check ./cmd/server

# -----------------------------
# Etapa 2: Imagen final
# -----------------------------
FROM alpine:3.20

WORKDIR /app

# Copiar el binario compilado al build
COPY --from=builder /app/health-check .

# Copiar el .env en la configuracion de entorno dentro del contenedor
COPY .env .env

# Exponer el puerto
EXPOSE 8081

# Correr el servicio (binario)
CMD ["./health-check"]
