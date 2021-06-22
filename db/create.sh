#!/bin/bash
#
#
. config.sh

docker build -t "mongodb-${APP}" .
