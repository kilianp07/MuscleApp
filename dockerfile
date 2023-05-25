From golang:1.18

WORKDIR /go/src/app

COPY . .

RUN go build -o api main.go

CMD ["./api"]
