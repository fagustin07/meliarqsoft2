# imagen base
FROM golang:1.20.3-alpine3.17

# directorio de trabajo
WORKDIR /app

# copiar archivos necesarios
COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
COPY docs ./docs
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# copiar código fuente
COPY cmd cmd/

# compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/

# imagen final
FROM golang:1.20.3-alpine3.17
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY .env .
COPY --from=0 /app .

# exponer el puerto 8080
EXPOSE 8080

# comando de inicio
CMD ["./app"]
