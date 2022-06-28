FROM golang:1.18 AS build

RUN mkdir /source
COPY . /source
WORKDIR /source
RUN go env -w GOPROXY=https://goproxy.cn,direct && go env -w GOPRIVATE=.gitlab.com,.gitee.com && go env -w GOSUMDB="sum.golang.google.cn"
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM registry.cn-shanghai.aliyuncs.com/shysj/go-base:main as final
WORKDIR /service
COPY --from=build /source/app /service/app