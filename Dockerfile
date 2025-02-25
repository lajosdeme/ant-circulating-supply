FROM golang:1.23-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o supply

EXPOSE 3000

CMD ["./supply", "start"]