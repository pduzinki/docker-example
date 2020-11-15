FROM golang

RUN go get github.com/gorilla/mux
RUN go get github.com/jinzhu/gorm
RUN go get github.com/lib/pq

WORKDIR /go/src/go-in-docker-example
COPY . .

CMD ["go", "run", "main.go"]