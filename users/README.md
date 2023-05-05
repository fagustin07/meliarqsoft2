# User service

## Correr servicio
1) crear el archivo .env y copiar lo contenido dentro del .env.example, configurando la URI de la DB.
2) `docker build -t user_app .`
3) `docker run -p <LOCAL_PORT>:50051 user_app`
4) para ver la documentacion, entrar a `localhost:<LOCAL_PORT>/docs/index.html`

## Correr los tests
Primeramente, para correr los tests unitarios de la app, primero la debemos tener corriendo y ejecutar en el contenedor de la app:
1. `docker exec -it user_app sh -c 'go test -v -coverprofile="coverage.out" -covermode=atomic ./internal/domain/...'`
2. `docker exec -it user_app sh -c 'go install gitlab.com/fgmarand/gocoverstats@latest'`
3. `docker exec -it user_app sh -c 'gocoverstats -v -f coverage.out -percent > coverage_rates.out'`

Con estos pasos lo que se hizo fue descargar paquetes para correr y medir los tests con coverage(1,2) y generar un reporte de coverage(3). Por último, para visualizar el coverage de los tests corremos:

`docker exec -it user_app sh -c 'cat coverage_rates.out'`.
