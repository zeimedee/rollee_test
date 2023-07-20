FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o /service cmd/server/main.go

EXPOSE 4000

CMD ["/service"]