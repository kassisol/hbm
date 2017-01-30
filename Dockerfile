# docker build --rm --no-cache --build-arg version=x.x.x -t kassisol/hbm:x.x.x .
# docker run -t --rm -v /tmp/hbm:/tmp/hbm kassisol/hbm:x.x.x <all|binary|rpm>
#
FROM centos:7

MAINTAINER Julien K. <hbm@kassisol.com>

ARG version

ENV GO_VERSION 1.7.4

COPY . /go/src/github.com/kassisol/hbm/
WORKDIR /go/src/github.com/kasisol/hbm

RUN build="gcc git rpm-build" \
	&& set -x \
	&& yum -y install $build \
	# Move entrypoint.sh script to /
	&& mv /go/src/github.com/kassisol/hbm/entrypoint.sh /entrypoint.sh \
	# Go
	&& curl -SL https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz | tar -xzC /usr/local \
	&& mkdir -p /go/bin \
	&& mkdir /go/pkg \
	&& mkdir -p /go/src/github.com \
	&& PATH=$PATH:/usr/local/go/bin \
	&& export GOPATH=/go \
	# Build hbm binary
	&& go build -o /go/bin/hbm --ldflags "-X github.com/kassisol/hbm/version.Version=$version -linkmode external -extldflags -static" \
	&& go install \
	# Build RPM package
	&& sh hack/make.sh \
	# Remove Go
	&& unset GOPATH \
	&& rm -rf /usr/local/go \
	&& PATH=$(echo $PATH | sed -e 's;:\?/usr/local/go/bin;;') \
	# Uninstall $build packages
	&& yum -y remove $build \
	&& yum -y autoremove \
	&& yum clean all \
	# Remove Go SRC directory
	&& rm -rf /go/src

ENTRYPOINT ["/entrypoint.sh"]
