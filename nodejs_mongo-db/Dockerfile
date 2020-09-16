FROM node:14.10-alpine3.10

WORKDIR /usr/code/app/ 

COPY package.json .

RUN npm install 

COPY . . 

EXPOSE 5000

CMD "node" "server.js"