FROM wujason/alpine-cn
ADD api-gateway /data/api-gateway
ADD ./conf/*.ini /data/conf/
ADD resource/pri_key.pem resource/pub_key.pem /resource/
VOLUME /resource/
WORKDIR /data
ENTRYPOINT [ "/data/api-gateway" ]