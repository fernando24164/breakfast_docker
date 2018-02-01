FROM golang:1.9.3
MAINTAINER A3SEC

COPY rest.go /rest_app/rest.go

EXPOSE 8081

WORKDIR /rest_app/

RUN go build rest.go

CMD [ "./rest" ]