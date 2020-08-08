FROM golang:1.14
WORKDIR /go/src
COPY apiapp .
RUN go build .
RUN cp apiapp /opt
RUN chmod +x /opt/apiapp
RUN rm -rf /go/src
ENTRYPOINT ["/opt/apiapp"]
