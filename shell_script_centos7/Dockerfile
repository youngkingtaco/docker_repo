FROM centos:7

MAINTAINER rodri.ocalderon97@icloud.com

RUN yum update -y \
    yum install -y vim 

COPY processing.sh  /root/processing.sh

CMD ["/data/docker_repo/shell_script_centos7/processing.sh"]

RUN chmod +x /root/processing.sh

ENTRYPOINT ["/root/processing.sh"]
