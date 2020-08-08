FROM golang:1.14
WORKDIR /go/src
ADD webapp /go/src
ADD utils /go/src
WORKDIR /go/src/webapp
RUN go build .
RUN cp webapp /opt
RUN chmod +x /opt/app
RUN rm -rf /go/src
ENTRYPOINT ["/opt/app"]
