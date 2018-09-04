FROM golang

ADD . /go/src/beacon/hh-location

RUN go get -u github.com/go-sql-driver/mysql && \
    go get -u github.com/gorilla/mux && \
    go get -u github.com/jmoiron/sqlx && \
    go get -u github.com/tkanos/gonfig && \
    go get -u github.com/jinzhu/gorm && \
    go get -u github.com/qor/qor && \
    go get -u github.com/qor/admin

RUN go install beacon/hh-location

ENTRYPOINT /go/bin/hh-location

