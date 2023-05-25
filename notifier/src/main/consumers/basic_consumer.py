
from abc import ABC, abstractmethod


class BasicClient(ABC):

    @abstractmethod
    def callback(self, ch, method, properties, body):
        pass

    @abstractmethod
    def consume(self):
        pass

    @abstractmethod
    def close(self):
        pass
