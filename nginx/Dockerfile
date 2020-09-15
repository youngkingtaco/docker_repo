FROM nginx:1.19.2

MAINTAINER Hugo Calderon rodri_ocalderon97@icloud.com

RUN  apt-get update \
     && apt-get install -y htop \
     && apt-get install -y apt-utils 

EXPOSE 8080/tcp 

COPY index.html  /usr/share/nginx/html/

WORKDIR /usr/share/nginx/html 

HEALTHCHECK  --interval=5m  --timeout=15s \
CMD curl -f http://localhost/ || exit 1