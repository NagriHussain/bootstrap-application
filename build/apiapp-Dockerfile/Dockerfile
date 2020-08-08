FROM golang:1.14
WORKDIR /go/src
COPY apiapp /go/src/apiapp
COPY utils /go/src/utils
WORKDIR /go/src/apiapp
RUN go build .
RUN cp apiapp /opt
RUN chmod +x /opt/apiapp
RUN rm -rf /go/src
ENTRYPOINT ["/opt/apiapp"]
