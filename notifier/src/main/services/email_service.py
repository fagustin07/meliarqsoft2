from services.basic_email_service import BasicEmailService
from src.main.config.smtp import email, password
from email.message import EmailMessage
import smtplib


class EmailService(BasicEmailService):
    def __init__(self):
        self.smtp = smtplib.SMTP("smtp-mail.outlook.com", port=587)
        self.smtp.starttls()
        self.smtp.login(email, password)

    def send_email(self, to, body, title):
        msg = EmailMessage()
        msg['From'] = email
        msg['To'] = to
        msg['Subject'] = title
        msg.set_content(body, subtype='html')

        self.smtp.sendmail(from_addr=email, to_addrs=to, msg=msg.as_string())
