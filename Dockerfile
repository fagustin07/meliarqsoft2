# imagen base
FROM golang:1.17

# directorio de trabajo
WORKDIR /app

# copiar archivos necesarios
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./config ./config
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# copiar código fuente
COPY cmd/ cmd/

# compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/

# imagen final
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/app .

# exponer el puerto 8080
EXPOSE 8080

# comando de inicio
CMD ["./app"]
