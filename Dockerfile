FROM golang:alpine3.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# RUN go run main.go migrate

RUN go build -o main .

EXPOSE 1111

CMD ["./main", "httpsrv"]