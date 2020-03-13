#!/bin/bash
set -e

CVER=$(git tag --sort=-version:refname | head -n 1)
SHA=$(git rev-parse HEAD)

docker build \
        --build-arg SHA=${SHA} \
		--build-arg CVER=${CVER} \
        -t anz-test .