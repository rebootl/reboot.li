FROM node:16

WORKDIR /usr/src/app

# install npm packages first to make use of caching
COPY package*.json ./
RUN npm i

# copy files
COPY src ./src
COPY static ./static
COPY .env ./
COPY svelte.config.js ./
COPY vite.config.ts ./
#COPY tsconfig.json ./
#COPY config_prod.js ./config.js

# build
RUN npm run build

EXPOSE 3000

# start
CMD ["node", "build"]
