FROM nginx:1.19.2

LABEL "Contact Info"="rodri.ocalderone97@icloud.com"
LABEL "nginx=1.19.2"

RUN  apt-get update || apt-get install  -y ping 


#This will be executed as /bin/sh -c
CMD echo "This is a test"

#you need to expose the port when you run the docker container 
EXPOSE 8080/tcp 

COPY index.html  /usr/share/nginx/html/

WORKDIR /usr/share/nginx/html 

CMD cat /usr/share/nginx/html/index.html 

#This would a PID that will run with your create a docker container 
ENTRYPOINT ["top", "-b"]

HEALTHCHECK  --interval=5m  --timeout=3s \
CMD curl -f http://localhost/ || exit 1


