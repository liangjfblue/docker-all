FROM golang:1.13 AS build
RUN mkdir /build
WORKDIR	/build
ADD . .
WORKDIR	/build/cmd
ENV GO11MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM scratch AS prod
COPY --from=build /build/cmd/ .
CMD ["./cmd"]

