FROM wujason/alpine-cn
ADD exchange-srv /data/exchange-srv
ADD ./conf/*.ini /data/conf/
WORKDIR /data/
ENTRYPOINT [ "/data/exchange-srv" ]