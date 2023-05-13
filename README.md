# Arquitectura Hexagonal basada en DDD.

## Documentacion requerida
   Los documentos solicitados se encuentran en la carpeta *[documentacion](https://github.com/fagustin07/meliarqsoft2/tree/main/documentacion)*.
   
## Requerimientos
- [Docker](https://docs.docker.com/get-docker/). 🐳
- [Docker Compose](https://docs.docker.com/get-docker/). 🐳

## Deseable
- [Editor de Texto](http://territoriogo.blogspot.com/2018/10/que-editor-utilizar-para-programar-en-go.html). 📝

# Correr la app.

Para correr nuestro sistema, simplemente debemos ejecutar `docker compose --profile service up`

## Testing Unitario
Para correr tests unitarios, en cada servicio se encuentran los pasos correspondientes.

## Testing de Integracion
Para correr los tests de integracion debemos ejecutar:
1. `docker compose --profile service up`
2. Luego de que todos nuestros servicios se levanten, corremos: `docker compose --profile testing up`
3. Con esto veremos un resumen del reporte de los tests.


Autores: Mauro Bailon, Federico Sandoval.
