
FROM golang:1.16.5 AS builder
WORKDIR /opt
COPY . /opt
ARG version
ENV VERSION=$version
ARG full_version
ENV FULL_VERSION=$full_version
ENV GOPROXY="https://goproxy.cn,direct"

RUN go env -w GOPROXY=https://goproxy.cn,direct 
RUN go mod download \
    && make

FROM alpine:latest  
WORKDIR /root/
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g'  /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=builder /opt/gokit-service .
ENTRYPOINT ["./gokit-service"]  

# sample ussage :
# build: docker build -f ./Dockerfile -t gokit-service:0.0.1 .
# run : docker run -p 8888:8888 -p 6666:6666 -d --name gokit-service gokit-service:0.0.1

