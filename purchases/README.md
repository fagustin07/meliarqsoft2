# Purchase Service

## Correr servicio
1) Pararnos sobre `purchases/`.
2) crear el archivo .env y copiar lo contenido dentro del .env.example, configurando la URI de la DB y el servidor de rabbit provisto por los alumnos o uno propio.
3) `docker build -f DockerfileStandalone -t purchase_app .`
4) `docker run -p 50053:50053 --name purchase_app purchase_app`
5) para ver la documentacion, entrar a `localhost:<LOCAL_PORT>/docs/index.html`

## Correr los tests
Primeramente, para correr los tests unitarios de la app, primero la debemos tener corriendo y ejecutar en el contenedor de la app:
1. `docker exec -it purchase_app sh -c 'go test -v -coverprofile="coverage.out" -covermode=atomic ./internal/domain/...'`
2. `docker exec -it purchase_app sh -c 'go install gitlab.com/fgmarand/gocoverstats@latest'`
3. `docker exec -it purchase_app sh -c 'gocoverstats -v -f coverage.out -percent > coverage_rates.out'`

Con estos pasos lo que se hizo fue descargar paquetes para correr y medir los tests con coverage(1,2) y generar un reporte de coverage(3). Por último, para visualizar el coverage de los tests corremos:

`docker exec -it purchase_app sh -c 'cat coverage_rates.out'`.
