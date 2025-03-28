#!/bin/sh

# To use openapi-generator-cli, first install it via NPM:
#
#   npm install -g @openapitools/openapi-generator-cli
#
# You may also need to install java:
#
#   sudo apt install openjdk-21-jre-headless
#

export GO_POST_PROCESS_FILE="$(which gofmt) -w"

openapi-generator-cli generate -g go -i ../1click-api/openapi.yaml --enable-post-process-file \
  --git-user-id defuse-protocol --git-repo-id one-click-sdk-go --package-name oneclick