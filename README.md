# Arquitectura Hexagonal basada en DDD.

## Documentacion requerida
   Los documentos solicitados se encuentran en la carpeta *[documentacion](https://github.com/fagustin07/meliarqsoft2/tree/main/documentacion)*.
   
## Requerimientos
- [Docker](https://docs.docker.com/get-docker/). 🐳
- [Docker Compose](https://docs.docker.com/get-docker/). 🐳

## Deseable
- [Editor de Texto](http://territoriogo.blogspot.com/2018/10/que-editor-utilizar-para-programar-en-go.html). 📝

# Correr la app.
1. Copiar el contenido del archivo .env.example en un archivo .env
2. Solicitar las credenciales necesarias para que los servicios puedan funcionar correctamente.
3. Una vez configurado el archivo .env, simplemente debemos ejecutar `docker compose up`.
4. En nuestro navegador ingresamos a http://localhost:9000/docs/index.html, donde veremos los protocolos de comunicacion de nuestro sistema realizado con Swagger.

## Testing Unitario
Para correr tests unitarios, en cada servicio se encuentran los pasos correspondientes. Los servicios que poseen tests son: products, users y purchases.

## Testing de Integracion
Para correr los tests de integracion debemos ejecutar:
1. `docker compose --profile service -f docker-compose-test.yml up`
2. Luego de que todos nuestros servicios se levanten, corremos: `docker compose --profile testing -f docker-compose-test.yml up`
3. Con esto veremos un resumen del reporte de los tests.
Posible error y solucion: Hemos detectado que ciertas veces `docker compose` suele fallar al utilizar las mismas imagenes. Para solucionarlo, ejecute el comando en cuestion agregándole el argumento `--force-recreate` al final del mismo.

Autores: Mauro Bailon, Federico Sandoval.
