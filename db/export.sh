#!/bin/bash
#
#
. config.sh

for c in "${COLLECTIONS[@]}"; do
  docker exec -it "mongodb-${APP}" mongoexport \
    --authenticationDatabase=admin \
    -u "$USER" -p "$PASSWORD" \
    -d "$APP" -c "$c" \
    --jsonArray --pretty > "export/${c}.json"

done
