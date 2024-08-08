FROM node

WORKDIR /app

COPY package*.json /app

RUN npm install

COPY index.js /app/index.js
COPY deployment.mustache /app/deployment.mustache

RUN mkdir /mnt/website

CMD ["node", "/app/index.js"]
