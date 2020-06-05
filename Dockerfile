FROM golang:latest

WORKDIR /go/src/github.com/appointment
COPY . .
RUN go get golang.org/x/net/netutil
RUN go get github.com/go-openapi/strfmt
RUN go get github.com/go-openapi/swag
RUN go get github.com/go-openapi/validate
RUN go get github.com/docker/go-units
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jessevdk/go-flags

RUN go build ./http/cmd/healthy-calendar-server/main.go
CMD ["./main",  "--host", "0.0.0.0", "--port", "8080"]
