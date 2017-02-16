# docker build --rm --no-cache --build-arg version=x.x.x -t kassisol/hbm-binary:x.x.x .
# docker run -t --rm -v /tmp/hbm:/tmp/hbm kassisol/hbm-binary:x.x.x
#
FROM debian:jessie

MAINTAINER Julien K. <hbm@kassisol.com>

ARG version

ENV GO_VERSION 1.7.4

COPY . /go/src/github.com/kassisol/hbm
COPY hack/packages/binary/entrypoint.sh /entrypoint.sh

WORKDIR /go/src/github.com/kassisol/hbm

RUN echo 'debconf debconf/frontend select Noninteractive' | debconf-set-selections \
	&& build="ca-certificates gcc libc6-dev git curl" \
	&& set -x \
	&& apt-get update \
        && apt-get install -y --no-install-recommends \
                $build \
	# Go
	&& curl -SL https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz | tar -xzC /usr/local \
	&& mkdir -p /go/bin \
	&& mkdir /go/pkg \
	&& PATH=$PATH:/usr/local/go/bin \
	&& export GOPATH=/go \
	# Build hbm binary
	&& go build -o /go/bin/hbm --ldflags "-X github.com/kassisol/hbm/version.Version=$version -linkmode external -extldflags -static" \
	&& tar czf /usr/local/src/hbm-${version}-linux-amd64.tar.gz -C /go/bin hbm \
	# Remove Go
	&& unset GOPATH \
	&& rm -rf /usr/local/go \
	&& PATH=$(echo $PATH | sed -e 's;:\?/usr/local/go/bin;;') \
	# Remove Go SRC directory
	&& rm -rf /go/src \
	# Uninstall $build packages
	&& apt-get remove -y $build \
        && apt-get -y autoremove \
        && apt-get clean \
        && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ENTRYPOINT ["/entrypoint.sh"]
