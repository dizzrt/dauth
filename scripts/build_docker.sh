#!/bin/sh

# exit when any command fails
set -e

cd "$(dirname "$0")/../"
. ./scripts/build_init.sh

echo "Start building docker image dauth:${_VERSION}..."

docker buildx build -f ./scripts/Dockerfile \
    --memory=8g \
    --memory-swap=8g \
    --platform="${PLATFORM}" \
    --build-arg GOOS="${GOOS}" \
    --build-arg GOARCH="${GOARCH}" \
    --build-arg ENV="${ENV}" \
    --build-arg VERSION="${_VERSION}" \
    --build-arg GIT_COMMIT="${GIT_COMMIT}" \
    -t dauth/dauth:${_VERSION} .

echo "${GREEN}Completed building docker image ${_VERSION}.${NC}"
echo ""
echo "Command to start dauth"
echo ""
echo "$ docker run -d -p 8081:8081 -p 50051:50051 -v ~/.dauth/logs:/app/logs --name dauth dauth/dauth:${_VERSION}"
echo ""