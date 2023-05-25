import codecs

from src.main.consumers.notification_consumer import NotificationConsumer
from src.main.services.email_service import EmailService

if __name__ == '__main__':
    # Services
    email_service = EmailService()

    # Consumer
    consumer = NotificationConsumer(email_service)

    consumer.consume()

