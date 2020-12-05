FROM golang

RUN go get github.com/gorilla/mux \
           github.com/jinzhu/gorm \
           github.com/lib/pq

WORKDIR /go/src/docker-example
COPY . .

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

CMD /wait && go run main.go