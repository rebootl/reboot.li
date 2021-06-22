#!/bin/bash
#
#
. config.sh

docker exec "mongodb-${APP}" \
  sh -c 'exec mongodump --authenticationDatabase=admin \
    -u "$USER" -p "$PASSWORD" \
    -d "$APP" --archive' \
  > "dump/${APP}.dump"
