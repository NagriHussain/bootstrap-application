FROM golang:1.14
WORKDIR /go/src
COPY webapp .
RUN go build .
RUN cp webapp /opt
RUN chmod +x /opt/app
RUN rm -rf /go/src
ENTRYPOINT ["/opt/app"]
