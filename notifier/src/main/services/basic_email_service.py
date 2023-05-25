from abc import ABC, abstractmethod


class BasicEmailService(ABC):

    @abstractmethod
    def send_email(self, to, body, title):
        pass
