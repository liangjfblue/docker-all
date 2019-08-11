FROM golang:1.12 AS build

RUN mkdir -p /go/src/gin-docker
WORKDIR	/go/src/gin-docker

ADD ./gin-docker/ .

ENV GOPATH=/go
ENV GOROOT=/usr/local/go

ENV GO111MODULE=on
ENV GOPROXY=http://mirrors.aliyun.com/goproxy/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM scratch AS prod

COPY --from=build /go/src/gin-docker/gin-docker .
CMD ["./gin-docker"]

