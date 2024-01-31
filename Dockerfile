FROM golang:1.21

WORKDIR /go/src/delivery

COPY . .

ARG SERVER_ADDRESS=localhost
ARG SERVER_PORT=3000
ARG DB_USER=postgres
ARG DB_PASS=postgres
ARG DB_ADDR=localhost
ARG DB_PORT=3306
ARG DB_NAME=go-delivery

RUN go build -o delivery-api -ldflags="-X main.serverAddress=${SERVER_ADDRESS} -X main.serverPort=${SERVER_PORT} -X main.dbUser=${DB_USER} -X main.dbPass=${DB_PASS} -X main.dbAddr=${DB_ADDR} -X main.dbPort=${DB_PORT} -X main.dbName=${DB_NAME}"

EXPOSE ${SERVER_PORT}

CMD ["./delivery-api"]
