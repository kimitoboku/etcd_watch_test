#!/bin/sh
IMG=watcher-client:develop
docker build --file=cmd/watcher_client/docker/Dockerfile -t $IMG .
