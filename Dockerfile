FROM golang:1.22 as builder

WORKDIR /app

ADD . /app

RUN go build -o bin/server server/main.go

FROM ubuntu:latest

COPY --from=builder /app/bin/server .

EXPOSE 8080

CMD ["./server"]