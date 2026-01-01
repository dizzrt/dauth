#!/bin/sh
# prepare common variables & functions for the build scripts.

# exit when any command fails
set -e

# Global variables
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# for docker buildx
PLATFORM="linux/arm64"

ENV="prod"
_VERSION=""
GIT_COMMIT=$(git rev-parse HEAD)

# Try to get version from command line arguments
for arg in "$@"; do
    case $arg in
        --platform=*)
            PLATFORM="${arg#*=}"
            shift
            ;;
        --env=*)
            ENV="${arg#*=}"
            shift
            ;;
        --version=*)
            _VERSION="${arg#*=}"
            shift
            ;;
    esac
done

# If version not provided via command line, try to get from git tag
if [ -z "$_VERSION" ]; then
    _VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "")
fi

# If version not provided via git tag, try to get from environment variable
if [ -z "$_VERSION" ]; then
    _VERSION=${VERSION:-}
fi

# If version not provided via environment variable, set default version
if [ -z "$_VERSION" ]; then
    _VERSION="${ENV}_$(date +"%Y%m%d%H%M%S")"
fi

# Version function used for version string comparison
version() { echo "$@" | awk -F. '{ printf("%d%03d%03d%03d\n", $1,$2,$3,$4); }'; }

mkdir_output() {
    if [ -z "$1" ]; then
        mkdir -p output
        OUTPUT_DIR=$(cd output > /dev/null && pwd)
    else
        OUTPUT_DIR="$1"
    fi
    echo "$OUTPUT_DIR"
}

# Go and node version checks.
TARGET_GO_VERSION="1.25.5"
GO_VERSION=`go version | { read _ _ v _; echo ${v#go}; }`
if [ "$(version ${GO_VERSION})" -lt "$(version $TARGET_GO_VERSION)" ];
then
   echo "${RED}Precheck failed.${NC} Require go version >= $TARGET_GO_VERSION. Current version ${GO_VERSION}."; exit 1;
fi
