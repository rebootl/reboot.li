FROM node:14

WORKDIR /usr/src/app

# install npm packages first to make use of caching
COPY package*.json ./
COPY .npmrc ./
RUN npm i

# copy files
COPY svelte.config.js ./
COPY tsconfig.json ./

COPY static ./static
COPY src ./src

COPY config.js_prod ./config.js

# build
RUN  npm run build

EXPOSE 3000

# start
CMD ["node", "build"]
