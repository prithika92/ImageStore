FROM golang:1.8 as goimage
ENV SRC=/go/src/
RUN mkdir -p /go/src/
RUN mkdir /go/src/go_docker
WORKDIR /go/src/go_docker
RUN cd /go/src/go_docker
COPY StoreImage.go .
RUN go build -o /go/bin/go_docker
ENTRYPOINT /go/bin/go_docker
EXPOSE 8080