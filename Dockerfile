FROM golang:1.19-alpine

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

RUN go build -o /app/default-http-backend

EXPOSE 8080

CMD [ "/app/default-http-backend" ]
