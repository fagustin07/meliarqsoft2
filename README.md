# Arquitectura Hexagonal basada en DDD.

## Requerimientos
- [Docker](https://docs.docker.com/get-docker/). 🐳
- [Docker Compose](https://docs.docker.com/get-docker/). 🐳

## Deseable
- [Editor de Texto](http://territoriogo.blogspot.com/2018/10/que-editor-utilizar-para-programar-en-go.html). 📝

# Correr la app.
1. Clonamos el proyecto  `git clone https://github.com/fagustin07/meliarqsoft2.git`.
2. Creamos un archivo `.env` y copiamos el contenido del archivo dado `.env.example`.
3. En este punto tenemos que correr la app, teniendo las siguientes consideraciones: 
   1. Si queremos correr todo a traves de docker compose, pasamos al siguiente punto.
   2. Si queremos correr la app con la DB de Mongo Atlas, debemos solicitar al autor del proyecto y reemplazar la key `MONGODB_URI`.
   3. Si queremos correr solo la app (implica instalar Golang en nuestra PC) y dockerizar la db, debemos reemplazar en `MONGODB_URI` el host `db` por `localhost`.
4. Con base en lo que hayamos elegido, seguimos de forma correlativa:
   1. Para correr todo el compose local, ejecutamos `docker compose -f docker-compose-dev.yml up`
   2. Para correr la app con Mongo Atlas, ejecutamos `docker compose up`
   3. Para correr solo la DB con la app local:
      1. Ejecutamos `docker compose -f docker-compose-dev-db.yml`.
      2. Ejecutamos `go run cmd/main.go`
5. Para este punto, deberiamos ver en la consola un mensaje que nos indique que la app corre en el puerto 8080.
6. Entramos a nuestro navegador y buscamos `localhost:8080/docs/index.html` y veremos la doc heca con Swagger de la app.

# Correr los tests
Primeramente, para correr los tests unitarios de la app, primero la debemos tener corriendo y ejecutar en el contenedor de la app:
1. `docker exec -it <CONTENEDOR> sh -c 'go test -v -coverprofile="coverage.out" -covermode=atomic ./internal/domain/...'`
2. `docker exec -it <CONTENEDOR> sh -c 'go install gitlab.com/fgmarand/gocoverstats@latest'`
3. `docker exec -it <CONTENEDOR> sh -c 'gocoverstats -v -f coverage.out -percent > coverage_rates.out'`
Con estos pasos lo que se hizo fue descargar paquetes para correr y medir los tests con coverage(1,2) y generar un reporte de coverage(3). Por último, para visualizar el coverage de los tests corremos:
4. `docker exec -it <CONTENEDOR> sh -c 'cat coverage_rates.out'`.

Aclaracion: Por defecto el nombre del contenedor de la app es meliarq-app-1. Puede asegurarse con el comando `docker ps` el nombre del contenedor es el basado en la imagen `meliarq-app`

### Autor: Federico Sandoval.