===============
Working
===================
FROM golang:1.8 as goimage
ENV SRC=/go/src/
RUN mkdir -p /go/src/
WORKDIR /go/src/go_docker
COPY StoreImage.go /go/src/go_docker/ \
&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o bin/go_docker

======
Working_final
=======
FROM golang:1.8 as goimage
ENV SRC=/go/src/
RUN mkdir -p /go/src/
RUN mkdir /go/src/go_docker
WORKDIR /go/src/go_docker
RUN cd /go/src/go_docker
COPY StoreImage.go .
RUN go build -o /bin/go_docker




FROM alpine:3.8 as baseimagealp
# Environment source
ENV SRC=go/src/
# create a working directory
RUN mkdir -p go/src/store_image
WORKDIR go/src/
ENTRYPOINT go/src/store_image
# run StoreImage.go
CMD ["go", "run", "src/store_image/StoreImage.go"]
EXPOSE 8080


FROM golang:1.8 as goimage
ENV SRC=/go/src/
RUN mkdir -p /go/src/
WORKDIR /go/src/go_docker
RUN git clone -b <Specify-branchname> — single-branch <Github HTTP Url> /go/src/go_docker/ \
&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o bin/go_docker


FROM alpine:3.6 as baseimagealp
RUN apk add--no-cache bash
ENV WORK_DIR=/docker/bin
WORKDIR $WORK_DIR
COPY--from=goimage /go/src/go_docker/bin/ ./
ENTRYPOINT /docker/bin/go_docker
EXPOSE 8080


FROM golang:1.8 as goimage
ENV SRC=/go/src/
RUN mkdir -p /go/src/
WORKDIR /go/src/go_docker
COPY . /go/src/go_docker/ 
go build -o bin/go_docker