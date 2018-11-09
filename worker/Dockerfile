FROM golang:1.10
USER root

# Dependencies
RUN apt-get update
RUN apt-get install -y uuid uuid-dev uuid-runtime
#RUN apt-get install -y libzmq-dbg libzmq-dev libzmq1
RUN apt-get install -y build-essential gcc wget
RUN apt-get install -y pkg-config

# ZeroMQ library
WORKDIR /tmp
RUN wget https://github.com/zeromq/libzmq/releases/download/v4.2.5/zeromq-4.2.5.tar.gz
RUN tar xvf zeromq-4.2.5.tar.gz

WORKDIR zeromq-4.2.5
RUN ./configure
RUN make install

USER root
RUN make install
RUN ldconfig

RUN go get github.com/pebbe/zmq4