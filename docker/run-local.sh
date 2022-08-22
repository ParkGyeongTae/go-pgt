#!/bin/bash

source .env

CONTAINER_NAME=${GO_CONTAINER_NAME}

IMAGE_NAME_TAG=$GO_IMAGE_NAME:$GO_IMAGE_VERSION

if [[ "$(docker images -f "dangling=true" -q 2> /dev/null)" != "" ]]; then
    echo "Remove dangling images"
    docker rmi $(docker images -f "dangling=true" -q)
fi

if [[ "$(docker container ls -q --filter "name=${CONTAINER_NAME}" 2> /dev/null)" != "" ]]; then
    echo "Stop & Remove Container ${CONTAINER_NAME}"
    docker stop $(docker container ls -q --filter "name=${CONTAINER_NAME}")
    docker rm $(docker container ls -q --filter "name=${CONTAINER_NAME}")
fi

docker run -it --rm \
    -v $(pwd)/code:/code \
    -v $(pwd)/result:/result \
    -p 9000:9000 \
    --name ${CONTAINER_NAME} ${IMAGE_NAME_TAG}