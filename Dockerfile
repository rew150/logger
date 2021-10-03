FROM golang:1.17

WORKDIR /app

COPY . .

RUN make build

EXPOSE 8080

CMD make start-prod
