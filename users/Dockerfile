# imagen base
FROM golang:1.20.3-alpine3.17

# directorio de trabajo
WORKDIR /app

# copiar archivos necesarios
COPY users/cmd ./cmd
COPY users/internal ./internal
COPY users/pkg ./pkg
COPY users/docs ./docs
COPY users/go.mod ./
COPY users/go.sum ./

RUN go mod download

# copiar código fuente
COPY users/cmd cmd/

# compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/

# imagen final
FROM golang:1.20.3-alpine3.17
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY .env .
COPY --from=0 /app .

# comando de inicio
CMD ["./app"]
