FROM golang:latest

ARG USER_ID
ARG USERNAME
ARG GROUP_ID

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

WORKDIR /app

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

RUN addgroup --gid $USER_ID $USERNAME
RUN adduser --disabled-password --gecos '' --uid $USER_ID --gid $GROUP_ID $USERNAME
USER $USERNAME

CMD air