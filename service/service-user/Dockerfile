FROM wujason/alpine-cn
ADD user-srv /data/user-srv
ADD ./conf/*.ini /data/conf/
WORKDIR /data
ENTRYPOINT [ "/data/user-srv" ]