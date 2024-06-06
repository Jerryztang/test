FROM csighub.tencentyun.com/admin/tlinux2.2-bridge-tcloud-underlay:latest

# 维护者信息
LABEL maintainer="jerryztang@tencent.com"

# 定义Go版本和安装路径
ENV GOLANG_VERSION 1.20.10
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

# 安装Go 以及 铁将军更新升级
RUN yum install -y wget && \
    wget https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    rm go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    go install github.com/erning/gorun@latest && \
    mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH" && \
    cd /tmp && wget https://mirrors.tencent.com/repository/generic/TSC/tsc_agent/allin1tscagent.tar.gz && tar zxf allin1tscagent.tar.gz && cd allin1tscagent && sh install.sh && \
    cd /tmp && wget http://tjj.woa.com/download/iGeneral_client_3.tgz && tar zxf iGeneral_client_3.tgz && cd iGeneral_client_3 && sh install.sh

