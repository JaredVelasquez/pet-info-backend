# Usa la imagen oficial de Go 1.24-alpine como imagen base
FROM golang:1.24-alpine

# Instala las herramientas necesarias
RUN apk add --no-cache git curl

# Descargar el binario precompilado de air
RUN curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia el archivo de configuración de air y el código fuente
COPY . .

# Comando por defecto para ejecutar air
CMD ["air"]
