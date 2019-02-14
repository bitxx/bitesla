FROM wujason/alpine-cn
ADD api-srv /data/api-srv
ADD ./conf/*.ini /data/conf/
ADD resource/pri_key.pem resource/pub_key.pem /resource/
VOLUME /resource/
WORKDIR /data
ENTRYPOINT [ "/data/api-srv" ]