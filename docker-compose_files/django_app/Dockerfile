FROM python:3.7-alpine 

MAINTAINER rodri.ocalderone97@icloud.com  Hugo Calderon 

RUN apk add --update --no-cache --virtual .tmp gcc libc-dev linux-headers 

COPY requirements.txt  /requirements.txt 

RUN pip install -r /requirements.txt 

RUN apk del .tmp 

WORKDIR /app 

COPY ./myapp /app

WORKDIR /scripts

COPY ./scripts /scripts 

RUN chmod +x /scripts/* 

RUN mkdir -p /vol/web/media 

RUN mkdir -p /vol/web/static 

RUN adduser -D user

RUN chown -R user:user /vol

RUN chmod -R 755 /vol/web

USER  user

CMD ["entrypoint.sh"]
