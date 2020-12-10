build:
	make -C cmd/watcher_client
	make -C cmd/put_client


build-docker-image:
	./cmd/watcher_client/docker/build.sh
	./cmd/put_client/docker/build.sh
