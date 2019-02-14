FROM wujason/alpine-cn
ADD strategy-srv /data/strategy-srv
ADD ./conf/*.ini /data/conf/
WORKDIR /data
ENTRYPOINT [ "/data/strategy-srv" ]