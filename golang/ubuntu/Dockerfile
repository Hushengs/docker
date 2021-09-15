FROM ubuntu:20.04
ENV PATH /usr/local/go/bin:$PATH
RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list \ 
	&& ln -snf /usr/share/zoneinfo/$TIMEZONE /etc/localtime && echo $TIMEZONE > /etc/timezone \
	&& apt update && apt install -y locales wget g++ gcc libc6-dev make pkg-config \
	&&	cd /usr/local \
	&&	wget https://golang.google.cn/dl/go1.17.1.linux-amd64.tar.gz \
	&& tar -C /user/local xf go1.17.1.linux-amd64.tar.gz \
	&& rm -rf go1.17.1.linux-amd64.tar.gz \
	&& go env -w GOPROXY=https://goproxy.cn,direct \
	&& go env -w GO111MODULE=on
WORKDIR /wwwroot