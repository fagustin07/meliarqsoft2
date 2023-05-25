# imagen base
FROM golang:1.20.3-alpine3.17

# directorio de trabajo
WORKDIR /app

# copiar archivos necesarios
COPY purchases/cmd ./cmd
COPY purchases/internal ./internal
COPY purchases/pkg ./pkg
COPY purchases/docs ./docs
COPY purchases/go.mod ./
COPY purchases/go.sum ./

RUN go mod download

# copiar código fuente
COPY purchases/cmd cmd/

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
