FROM alpine:3.8

ARG version

LABEL io.harbormaster.image.maintainer="hbm@kassisol.com"
LABEL io.harbormaster.image.version=$version
LABEL io.harbormaster.image.description="HBM is an application to authorize and manage authorized docker command using Docker AuthZ plugin"

COPY build/hbm /usr/local/sbin/hbm
COPY entrypoint.sh /entrypoint.sh

RUN apk --update --no-cache add bash

SHELL ["bash", "-c"]

RUN mkdir -p /etc/docker/plugins \
	&& mkdir -p /run/docker/plugins \
	&& echo -e "#!/bin/bash\\n\\ntrue" > /usr/local/bin/docker \
	&& chmod +x /usr/local/bin/docker

ENTRYPOINT ["/entrypoint.sh"]
