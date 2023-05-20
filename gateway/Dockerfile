# imagen base
FROM golang:1.20.3-alpine3.17

# directorio de trabajo
WORKDIR /app

# copiar archivos necesarios
COPY gateway/cmd ./cmd
COPY gateway/internal ./internal
COPY gateway/pkg ./pkg
COPY gateway/docs ./docs
COPY gateway/go.mod ./
COPY gateway/go.sum ./

RUN go mod download

# copiar código fuente
COPY gateway/cmd cmd/

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
