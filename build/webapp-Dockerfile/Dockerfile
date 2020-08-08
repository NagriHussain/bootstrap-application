FROM golang:1.14
WORKDIR /go/src
COPY webapp .
COPY utils .
RUN go build .
RUN cp webapp /opt
RUN chmod +x /opt/app
RUN rm -rf /go/src
ENTRYPOINT ["/opt/app"]
