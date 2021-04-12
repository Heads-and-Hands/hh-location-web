FROM golang

ADD . /go/src/hh-location-web/hh-location

RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/jmoiron/sqlx
RUN go get -u github.com/tkanos/gonfig
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/qor/qor
RUN go get -u github.com/qor/admin

RUN go install hh-location-web/hh-location

ENTRYPOINT /go/bin/hh-location

