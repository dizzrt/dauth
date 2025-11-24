#!/bin/sh

# exit when any command fails
set -e

cd "$(dirname "$0")/../"
. ./scripts/build_init.sh

echo "Start building docker image ${_VERSION}..."

docker buildx build -f ./scripts/Dockerfile \
    --memory=8g \
    --memory-swap=8g \
    --platform="${PLATFORM}" \
    --build-arg GOOS="${GOOS}" \
    --build-arg GOARCH="${GOARCH}" \
    --build-arg ENV="${ENV}" \
    --build-arg VERSION="${_VERSION}" \
    --build-arg GIT_COMMIT="${GIT_COMMIT}" \
    -t dauth/dauth .

echo "${GREEN}Completed building docker image ${_VERSION}.${NC}"
echo ""
echo "Command to tag and push the image"
echo ""
echo "$ docker tag dauth/dauth dauth/dauth:${_VERSION}; docker push dauth/dauth:${_VERSION}"
# echo ""
# echo "Command to start dauth on port 8080"
# echo ""
# echo "$ docker run --init --name dauth --restart always --publish 8080:8080 --volume ~/.dauth/data:/var/opt/dauth dauth/dauth:${_VERSION} --data /var/opt/dauth --port 8080"
# echo ""