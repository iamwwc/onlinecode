#/usr/bin/bash

docker login -u ${DOCKER_LOGIN_USER_NAME}  -p ${DOCKER_LOGIN_PASSWORD}
docker push "${DOCKER_REPO_NAME}/${IMAGE_NAME}:latest"
docker push "${DOCKER_REPO_NAME}/${IMAGE_NAME}:${TRAVIS_BUILD_NUMBER}"