FROM node:16

WORKDIR /app

COPY make-purchase-orchestator/package*.json ./

COPY make-purchase-orchestator ./
COPY .env .

RUN npm install

CMD ["node", "/app/src/app.js"]