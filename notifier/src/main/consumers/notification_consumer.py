import codecs
import json
from src.main.consumers.basic_consumer import BasicClient
from src.main.services.email_service import EmailService
from src.main.utils.content.sell_html import sell_html
from src.main.utils.content.purchase_html import purchase_html
from src.main.utils.content.register_html import register_html


class NotificationConsumer(BasicClient):

    def __init__(self, email_service: EmailService):
        self.email_service = email_service
        self.queue = 'notifier'
        self.channel.queue_declare(queue=self.queue)

    def callback(self, ch, method, properties, body):
        try:
            data = json.loads(body)
            type_msg = data['Key']
            to = data['Email']
            name = data['Name']

            # Notificacion de compra
            if type_msg == 'Sell':
                body = sell_html\
                    .replace("{user}", name)\
                    .replace("{product}", data['Product'])
                self.email_service.send_email(to=to, body=body, title="Notificacion de compra")

            # Notificacion de venta
            if type_msg == 'Purchase':
                body = purchase_html\
                    .replace("{user}", name)\
                    .replace("{product}", data['Product'])
                self.email_service.send_email(to=to, body=body, title="Compraste un producto")

            # Notificacion de Registro
            if type_msg == 'Register':
                body = register_html\
                    .replace("{user}", name)
                self.email_service.send_email(to=to, body=body, title="Bienvenido a Meli")
        except Exception as e:
            print(e)
