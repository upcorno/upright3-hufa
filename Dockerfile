FROM golang:alpine AS build

RUN mkdir /source
COPY . /source
WORKDIR /source
RUN go env -w GOPROXY=https://goproxy.cn,direct && go env -w GOPRIVATE=.gitlab.com,.gitee.com && go env -w GOSUMDB="sum.golang.google.cn"
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine:latest as final
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update
RUN apk add --no-cache tzdata
WORKDIR /service
COPY --from=build /source/app /service/app