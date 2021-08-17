#!/bin/bash

set -ex
IMAGE_NAME=$1
DOCKER_FILE="dummy-domain-service/Dockerfile"
GITHUB_SSH_PRIVATE_KEY="/home/go/.ssh/github_rsa"

docker version

echo Current Dir: "${PWD}"
echo Image Name: "${IMAGE_NAME}"

DOCKER_BUILDKIT=1 docker build --ssh github=${GITHUB_SSH_PRIVATE_KEY} -t ${IMAGE_NAME} -f ${DOCKER_FILE} .
docker push ${IMAGE_NAME}
