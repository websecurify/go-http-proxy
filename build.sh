#!/usr/bin/env bash

# ---
# ---
# ---

CSD=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

# ---
# ---
# ---

docker run \
	--rm \
	-v "${CSD}:/src" \
	-v "/var/run/docker.sock:/var/run/docker.sock" \
	"centurylink/golang-builder" \
	"websecurify/go-http-proxy"
	
# ---
# ---
# ---

rm "${CSD}/go-http-proxy"

# ---
