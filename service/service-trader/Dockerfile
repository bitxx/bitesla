FROM wujason/bitesla-alpine:latest
MAINTAINER  wj
ADD trader-srv /data/trader-srv
ADD ./conf/*.ini /data/conf/

WORKDIR /data
ENTRYPOINT [ "/data/trader-srv" ]