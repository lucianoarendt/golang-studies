FROM golang:1.17

WORKDIR /go/src/app

COPY main.go ./main.go
COPY go.mod ./go.mod

RUN go get -d ./
RUN go build -o main .

RUN adduser --disabled-password --gecos --quiet pyroscope
USER pyroscope

CMD ["./main"]