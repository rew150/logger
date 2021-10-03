FROM node:14-alpine AS front

WORKDIR /front

COPY frontend/package*.json ./

RUN npm install

COPY frontend/ ./

RUN npm run build

FROM golang:1.17

WORKDIR /app

COPY . .

RUN make build-back

COPY --from=front /front/public/ ./frontend/public/

EXPOSE 8080

CMD make start-prod
