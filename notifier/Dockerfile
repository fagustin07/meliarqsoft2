FROM python:3.9

WORKDIR /app

COPY notifier/ ./

RUN pip install --no-cache-dir -r requirements.txt

ENV PYTHONPATH /app

CMD ["python3", "./src/main/main.py"]