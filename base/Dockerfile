FROM golang:1.10

USER root

# Dependencies
RUN apt-get update
RUN apt-get install -y uuid uuid-dev uuid-runtime build-essential gcc wget pkg-config

# ZeroMQ library
WORKDIR /base
RUN wget https://github.com/zeromq/libzmq/releases/download/v4.2.5/zeromq-4.2.5.tar.gz
RUN tar xvf zeromq-4.2.5.tar.gz; rm zeromq-4.2.5.tar.gz

WORKDIR zeromq-4.2.5
RUN ./configure; make install; ldconfig

RUN go get github.com/pebbe/zmq4
