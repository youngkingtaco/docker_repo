FROM centos:latest

MAINTAINER hugo calderon rodri_ocalderon97@icloud.com
RUN yum update -y && \
    yum install -y vim 

COPY passwd.text / 

CMD ["cat","passwd.text"]
