FROM wujason/alpine-cn
ADD api-srv /data/api-srv
ADD ./conf/*.ini /data/conf/
VOLUME /resource/
WORKDIR /data
ENTRYPOINT [ "/data/api-srv" ]