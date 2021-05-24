# Please keep up to date with the new-version of Golang docker for builder
FROM golang:latest

ARG USER_ID
ARG GROUP_ID

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

WORKDIR /app

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

RUN addgroup --gid $USER_ID arichard
RUN adduser --disabled-password --gecos '' --uid $USER_ID --gid $GROUP_ID arichard
USER arichard

CMD air