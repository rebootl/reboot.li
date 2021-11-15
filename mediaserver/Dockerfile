FROM node:14

WORKDIR /usr/src/app

# install npm packages first to make use of caching
COPY package*.json ./
RUN npm i

# copy files
COPY db.js ./db.js
COPY imageStorage.js ./imageStorage.js
COPY main.js ./main.js

COPY config_prod.js ./config.js

RUN mkdir -p ./media

EXPOSE 3005

# start
CMD ["npm", "run", "start"]
