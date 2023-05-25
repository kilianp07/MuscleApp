FROM golang:1.18
LABEL org.opencontainers.image.source="https://github.com/kilianp07/MuscleApp"

WORKDIR /go/src/app

COPY . .

RUN go build -o api main.go

CMD ["./api"]
