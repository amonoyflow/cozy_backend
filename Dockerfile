FROM golang:latest

RUN mkdir /cozy

ADD . /cozy

WORKDIR /cozy

RUN go build -o main .

CMD [ "/cozy/main" ]