# Notifier Service

Implementación de un servicio de notificaciones para una plataforma de compra y venta de productos.

## Stack tecnologico

- Python
- [RabbitMQ](https://cloudamqp.com)
- Pika
- Docker

## Ejecucion del servicio

1. Configurar credenciales
	1. 	Credenciales para amqp
		Dentro del archivo [amqp.py](https://github.com/fagustin07/meliarqsoft2/blob/dev/notifier/src/main/config/amqp.py) completar los campos "user" y "password" con las credenciales de [cloudamqp](https://cloudamqp.com).
	2. Credenciales para envios de email
		Dentro del archivo [smtp.py](https://github.com/fagustin07/meliarqsoft2/blob/dev/notifier/src/main/config/smtp.py) completar los campos "email" y "password" con las credenciales de hotmail.
2. Para ejecutar el proyecto podemos hacerlo de forma normal o con docker:
   1. Para ejecutarlo de manera normal, primero debemos instalar las dependencias:

              Python3 –m pip install –r requirements.txt
   2. Finalmente ejecutamos la app con el siguiente comando:
   
              Python3 /src/main/main.py
   
   3. Para correr la app con docker primero generamos la imagen:

              docker build -t notifier-service .

   4. Por ultimo ejecutamos la imagen generada:

              docker run --name notifier-service notifier-service
   
