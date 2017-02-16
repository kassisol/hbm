# docker build --rm --no-cache --build-arg version=x.x.x -t kassisol/hbm:x.x.x .
# docker run -t --rm -v /tmp/hbm:/tmp/hbm kassisol/hbm:x.x.x <all|binary|rpm>
#
FROM centos:7

MAINTAINER Julien K. <hbm@kassisol.com>

COPY hbm.spec /usr/local/src/hbm.spec
COPY entrypoint.sh /entrypoint.sh
COPY hbm/build-rpm.sh /usr/local/bin/build-rpm.sh
COPY hbm/hbm.service /usr/local/src/hbm/hbm.service
COPY hbm/hbm-*-linux-amd64.tar.gz /tmp/hbm.tar.gz

RUN build="rpm-build" \
	&& set -x \
	&& yum -y install $build \
	# Build rpm
	&& /usr/local/bin/build-rpm.sh \
	# Uninstall $build packages
	&& yum -y remove $build \
	&& yum -y autoremove \
	&& yum clean all

ENTRYPOINT ["/entrypoint.sh"]
