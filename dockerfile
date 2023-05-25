From golang:1.18
LABEL org.opencontainers.image.source="https://github.com/kilianp07/MuscleApp"
WORKDIR /go/src/app

COPY . .

RUN go build -o main main.go

CMD ["./main"]
