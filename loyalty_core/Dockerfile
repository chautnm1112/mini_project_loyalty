FROM golang:1.23

RUN apt-get update && apt-get install -y postgresql-client

WORKDIR /app
COPY . .

RUN go build -o loyalty-core .

CMD ["./loyalty-core"]
