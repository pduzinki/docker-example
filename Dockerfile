FROM golang:latest as builder

WORKDIR /go/src/docker-example
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go/src/docker-example/app .

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

CMD /wait && ./app