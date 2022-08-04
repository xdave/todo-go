FROM golang:buster

RUN apt-get update && apt-get upgrade

RUN apt-get install git

RUN git clone https://github.com/brilliant-ember/todo-go

RUN cd todo-go

RUN go get github.com/google/uuid

