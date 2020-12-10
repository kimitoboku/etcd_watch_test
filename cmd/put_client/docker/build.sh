#!/bin/sh
IMG=put-client:develop
docker build --file=cmd/put_client/docker/Dockerfile -t $IMG .
