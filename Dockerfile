# Usar una imagen base de Golang con soporte para módulos
FROM golang:1.22 as builder

# Establecer el directorio de trabajo en el contenedor
WORKDIR /app

# Copiar los archivos del proyecto al contenedor
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Imagen mínima para ejecutar la aplicación (alpine)
FROM alpine:latest

WORKDIR /root/

# Copiar el binario compilado desde el builder
COPY --from=builder /app/main .

# Exponer el puerto del servidor
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
