FROM python:3.8-alpine 

MAINTAINER rodri.ocalderone97@icloud.com  Hugo Calderon 

ENV PATH="/scripts:${PATH}"

RUN apk add --update --no-cache --virtual .tmp gcc libc-dev linux-headers 

COPY requirements.txt  /requirements.txt 

RUN pip install -r /requirements.txt 

RUN apk del .tmp 

RUN mkdir /app 

COPY ./myapp /app

WORKDIR /app

COPY ./scripts /scripts 

RUN chmod +x /scripts/* 

RUN mkdir -p /vol/web/media 

RUN mkdir -p /vol/web/static 

RUN adduser -D user

RUN chown -R user:user /vol

RUN chmod -R 755 /vol/web

USER  user

WORKDIR /myapp

RUN chmod +x manage.py

CMD ["entrypoint.sh"]
