#!/bin/bash

docker run \
    --detach \
    --name=test_db\
    --env="MYSQL_ROOT_PASSWORD=root" \
    -p 3306:3306 \
    mysql