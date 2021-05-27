# Please keep up to date with the new-version of Golang docker for builder
FROM golang:latest

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

WORKDIR /app 

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air

## Production environment (alias: base)
#FROM golang:1.12-alpine as base
#RUN apk update && apk upgrade && \
#apk add --no-cache bash git openssh
#WORKDIR /home/my-project
#
## Development environment
## Unfortunately, linux alpine doesn't have fswatch package by default, so we will need to download source code and make it by outselves.
#FROM base as dev
#RUN apk add --no-cache autoconf automake libtool gettext gettext-dev make g++ texinfo curl
#WORKDIR /root
#RUN wget https://github.com/emcrisostomo/fswatch/releases/download/1.14.0/fswatch-1.14.0.tar.gz
#RUN tar -xvzf fswatch-1.14.0.tar.gz
#WORKDIR /root/fswatch-1.14.0
#RUN ./configure
#RUN make
#RUN make install
#WORKDIR /home/my-project