FROM golang:1.9.3
LABEL maintainer="fernando24164"

RUN mkdir /rest_app

COPY rest.go /rest_app

EXPOSE 8081

WORKDIR /rest_app

RUN go build rest.go

CMD [ "./rest" ]