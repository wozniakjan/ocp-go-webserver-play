FROM golang:1.10-alpine

USER nobody

RUN mkdir -p /go/src/github.com/pschiffe/ocp-go-webserver
WORKDIR /go/src/github.com/pschiffe/ocp-go-webserver

COPY . /go/src/github.com/pschiffe/ocp-go-webserver
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "main.go"]
