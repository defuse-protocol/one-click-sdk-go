#!/bin/bash
set -e

VERSION_NUMBER=${API_VERSION#v}
echo "Generating SDK for API version: $VERSION_NUMBER"

# Create temporary directory for working files
TEMP_DIR=$(mktemp -d)
trap 'rm -rf "$TEMP_DIR"' EXIT

# Download the OpenAPI specification
echo "Downloading OpenAPI specification..."
curl -s -o "$TEMP_DIR/openapi.yaml" "https://1click.chaindefuser.com/docs/v0/openapi.yaml"

# Install OpenAPI Generator CLI
npm install -g @openapitools/openapi-generator-cli

# Set up post-processing for Go files
export GO_POST_PROCESS_FILE="$(which gofmt) -w"

# Generate the SDK
echo "Generating Go SDK..."
openapi-generator-cli generate \
  -g go \
  -i "$TEMP_DIR/openapi.yaml" \
  --enable-post-process-file \
  --git-user-id defuse-protocol \
  --git-repo-id one-click-sdk-go \
  --package-name oneclick

# Commit and push changes
echo "Committing changes..."
git config user.email "action@github.com"
git config user.name "GitHub Action"
git add .
git commit -m "Release: ${VERSION_NUMBER} version under /v0 path" || echo "No changes to commit"

echo "Pushing changes..."
git tag -a "v$VERSION_NUMBER" -m "Release v$VERSION_NUMBER"
git push origin main
git push origin "v$VERSION_NUMBER"

echo "SDK regeneration and publishing completed successfully!"
