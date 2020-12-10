FROM golang:1.13 as builder
ADD . /go/src/github.com/kimitoboku/etcd_watch_test
RUN make -C /go/src/github.com/kimitoboku/etcd_watch_test/cmd/watcher_client

FROM centos:centos7
COPY --from=builder /go/src/github.com/kimitoboku/etcd_watch_test/cmd/watcher_client/watcher_client /usr/local/bin

