FROM node:9-slim as builder 

WORKDIR /app 

COPY package.json ./app 

RUN npm install 

COPY . /app

CMD ["npm","start"]
