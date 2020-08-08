FROM golang:1.14
WORKDIR /go/src
COPY webapp /go/src/webapp
COPY utils /go/src/utils
WORKDIR /go/src/webapp
RUN go build .
RUN cp webapp /opt
RUN chmod +x /opt/webapp
RUN rm -rf /go/src
ENTRYPOINT ["/opt/webapp"]
