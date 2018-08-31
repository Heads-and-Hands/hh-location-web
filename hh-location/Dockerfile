FROM golang

ADD . /go/src/beacon/hh-location

RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/jmoiron/sqlx
RUN go get -u github.com/tkanos/gonfig
RUN go get -u github.com/jinzhu/gorm

RUN go install beacon/hh-location

ENTRYPOINT /go/bin/hh-location

