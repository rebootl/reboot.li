FROM node:16

WORKDIR /usr/src/app

# install npm packages first to make use of caching
COPY package*.json ./
RUN npm i

# copy files
COPY db.js ./db.js
COPY imageStorage.js ./imageStorage.js
COPY main.js ./main.js

COPY config_prod.js ./config.js

# (media dir will be created autom.)
#RUN mkdir -p /usr/src/app/media

EXPOSE 3005

# start
CMD ["npm", "run", "start"]
