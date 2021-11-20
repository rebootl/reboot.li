FROM node:14

WORKDIR /usr/src/app

# install npm packages first to make use of caching
COPY package*.json ./
RUN npm i

# copy files
COPY svelte.config.js ./
COPY tsconfig.json ./

COPY static ./static
COPY src ./src

COPY config_prod.js ./config.js

# build
RUN npm run build

EXPOSE 3000

# start
CMD ["node", "build"]
