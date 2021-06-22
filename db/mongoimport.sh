#APP="better-reflectapp"
#USER="${APP}-admin"
#PASSWORD="example123"
#COLLECTIONS=( "users" "entries" )

. /docker-entrypoint-initdb.d/config.sh

for c in "${COLLECTIONS[@]}"; do
  mongoimport --db "$APP" \
    -u "${USER}" \
    -p "${PASSWORD}" \
    --authenticationDatabase admin \
    --collection "$c" \
    --jsonArray \
    --file "/docker-entrypoint-initdb.d/${c}.json"
done
