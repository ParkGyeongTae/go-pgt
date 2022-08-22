#!/bin/bash

source .env

docker rmi ${GO_IMAGE_NAME}:${GO_IMAGE_VERSION}