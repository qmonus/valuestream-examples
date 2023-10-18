#!/bin/bash

set -o xtrace

docker run \
    -p 8888:8888 \
    --env DB_HOST="host.docker.internal" \
    --env DB_PORT="3306" \
    --env DB_USER="root" \
    --env DB_PASS="example" \
    --env DB_NAME="database" \
    --env GIN_MODE="debug" \
    demo-backend-app
