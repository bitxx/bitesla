FROM golang:1.11-alpine
MAINTAINER  wj

RUN apk add --no-cache git

ADD trader-srv /data/trader-srv
ADD ./conf/*.ini /data/conf/

#WORKDIR /go/src/github.com/jason-wj/bitesla/
#COPY . .

WORKDIR /data
ENTRYPOINT [ "/data/trader-srv" ]