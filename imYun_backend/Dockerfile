FROM golang:1.14

# 镜像必要的环境变量
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn \
    GORUNENV=docker \
    GORUNADDR=0.0.0.0 \
    GORUNPORT=5000

WORKDIR /build

COPY . .

RUN go mod download && go build -o app main.go

WORKDIR /dist

RUN cp /build/app .
RUN mkdir /dist/UploadFile

EXPOSE 5000

CMD ["/dist/app"]
