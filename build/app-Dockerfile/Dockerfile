FROM golang:1.14
WORKDIR /go/src
COPY app .
RUN go build .
RUN cp app /opt
RUN chmod +x /opt/app
RUN rm -rf /go/src
ENTRYPOINT ["/opt/app"]
