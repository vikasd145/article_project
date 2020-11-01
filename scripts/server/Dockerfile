FROM golang:1.14

# Add Maintainer Info
LABEL maintainer="Vikas Dubey <vikas.dubey@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod vendor

COPY . .

RUN go build -v ./cmd/article_server/article_main.go -o main

EXPOSE 8080

CMD ["./main"]