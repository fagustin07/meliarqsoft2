import pika
from src.main.config.amqp import amqp_uri
from abc import ABC, abstractmethod


class BasicClient(ABC):

    params = pika.URLParameters(amqp_uri)
    connection = pika.BlockingConnection(params)
    channel = connection.channel()  # start a channel

    @abstractmethod
    def callback(self, ch, method, properties, body):
        pass

    def consume(self):
        self.channel.basic_consume(queue=self.queue, on_message_callback=self.callback, auto_ack=True)
        self.channel.start_consuming()

    def close(self):
        self.channel.close()
        self.connection.close()
