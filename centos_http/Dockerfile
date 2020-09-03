FROM centos:7

MAINTAINER rodri.ocalderone97@icloud.com

RUN yum update -y && \
    yum install -y httpd

EXPOSE 80 

COPY index.html /var/www/html/

ENTRYPOINT ["/usr/sbin/httpd"]

CMD ["-D","FOREGROUND"]

