FROM golang:1-alpine

ENV GIN_MODE=release
ENV PORT=8080

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /go/src/github.com/Chadiii/go-usermgmt-rest

ADD . .

RUN apk add curl
RUN apk add apache2-utils

RUN go build -o /go-usermgmt-rest

EXPOSE $PORT

CMD ["/go-usermgmt-rest"]