#!/bin/bash
#
#
. config.sh

docker stop "mongodb-${APP}"
docker rm "mongodb-${APP}"

sudo rm -rf mongodb-data/

./create.sh
./run.sh
