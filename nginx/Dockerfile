FROM nginx:1.19.2

LABEL "Contact Info"="rodri.ocalderone97@icloud.com"
LABEL "nginx"="1.19.2"

RUN  apt-get update \
     && apt-get install -y htop 

EXPOSE 8080/tcp 

COPY index.html  /usr/share/nginx/html/

WORKDIR /usr/share/nginx/html 



HEALTHCHECK  --interval=5m  --timeout=15s \
CMD curl -f http://localhost/ || exit 1
