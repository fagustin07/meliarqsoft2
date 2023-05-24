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
2. Solicitar las credenciales necesarias para que los servicios puedan funcionar correctamente con su base de datos.
3. Solicitar los archivos de configuracion para el notifier y reemplazar los existentes (ampq.py y smtp.py) en `notifier/src/config`
4. Una vez configurado el archivo .env, simplemente debemos ejecutar `docker compose up --build`.
5. En nuestro navegador ingresamos a http://localhost:9000/docs/index.html, donde veremos los protocolos de comunicacion de nuestro sistema realizado con Swagger.

## Testing Unitario
- Se desarrollaron tests unitarios para los servicios de: users, products y purchases.
- Para correr tests unitarios, en cada servicio mencionado se encuentran los pasos correspondientes.

## Testing de Integracion
Para correr los tests de integracion debemos ejecutar:
1. Si hemos corrido la app por si sola, primero ejecutamos `docker compose down`.
2. Copiar y pegar la información del archivo `.env.test.example` en nuestro `.env` y configurar el `AMPQ_URI` dado junto a `ENVIRONMENT` con el valor `test` (se recomienda guardar su .env actual para luego volver a correr la app).
3. A continuacion corremos `docker compose --profile service -f docker-compose-test.yml up --build`
4. Luego de que todos nuestros servicios se levanten, corremos en otra consola: `docker compose --profile tests -f docker-compose-test.yml up --build`
5. Con esto veremos un resumen del reporte de los tests.
6. Cada vez que queramos correr nuevamente los tests de integracion debemos
   - Ejecutar `docker compose --profile service -f docker-compose-test.yml down` para limpiar los contenedores de prueba.
   - Luego de que todos nuestros servicios se levanten, corremos en otra consola: `docker compose --profile tests -f docker-compose-test.yml up --force-recreate`

### Comentarios sobre test de integracion
- Posible error y solucion: Hemos detectado que ciertas al hacer el `docker compose ... up ...` tanto del perfil de servicio o del test, suele fallar al utilizar los mismos nombres de contenedores. Para solucionarlo, ejecute el comando en cuestion agregándole el argumento `--force-recreate` al final del mismo.
- Tener en cuenta que los entornos estan armados para correr estos tests una sola vez, por lo que importa el orden de la coleccion de postman y, a su vez, es importante volver a crear los contenedores para que se vuelvan a configurar en caso que se desee volver a correr los tests de integracion.



*Autores: Mauro Bailon, Federico Sandoval.*
