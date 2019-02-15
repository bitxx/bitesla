FROM wujason/alpine-cn
ADD trader-srv /data/trader-srv
ADD ./conf/*.ini /data/conf/
WORKDIR /data
ENTRYPOINT [ "/data/trader-srv" ]