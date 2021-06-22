#!/bin/bash
#
#
. config.sh

# options
#  mongod --replSet rs0

if [ "$PRODUCTION" = true ]; then
  docker run -d \
    -e MONGO_INITDB_ROOT_USERNAME="$USER" \
    -e MONGO_INITDB_ROOT_PASSWORD="$PASSWORD" \
    -e MONGO_INITDB_DATABASE="$APP" \
    -v "$(pwd)/mongodb-data:/data/db" \
    --restart unless-stopped \
    --name "mongodb-${APP}" \
    --network "${APP}-network" \
    "mongodb-${APP}" \
    mongod
else
  docker run -d -p "${LOCALPORT}":27017 \
    -e MONGO_INITDB_ROOT_USERNAME="$USER" \
    -e MONGO_INITDB_ROOT_PASSWORD="$PASSWORD" \
    -e MONGO_INITDB_DATABASE="$APP" \
    -v "$(pwd)/mongodb-data:/data/db" \
    --restart unless-stopped \
    --name "mongodb-${APP}" \
    "mongodb-${APP}" \
    mongod
fi
