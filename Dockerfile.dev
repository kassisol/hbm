FROM kassis/dev:go

########## Dapper Configuration #####################
ENV DAPPER_RUN_ARGS --privileged -p 127.0.0.1:2345:2345 --name hbm_dev
ENV DAPPER_SOURCE /go/src/github.com/kassisol/hbm
ENV SHELL /bin/bash

WORKDIR ${DAPPER_SOURCE}

########## General Configuration #####################
ARG DAPPER_HOST_ARCH=amd64
ARG HOST_ARCH=${DAPPER_HOST_ARCH}
ARG ARCH=${HOST_ARCH}

ARG APP_REPO=kassisol

ARG docker_version=18.03.0

# Set up environment and export all ARGS as ENV
ENV ARCH=${ARCH} \
	HOST_ARCH=${HOST_ARCH}

ENV DOCKER_COMPOSE_VERSION=1.14.0 \
	DAPPER_HOST_ARCH=${DAPPER_HOST_ARCH}

# Install Docker
RUN get-docker.sh $docker_version

# Install Docker Compose
RUN wget https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_VERSION}/docker-compose-Linux-x86_64 -O /usr/bin/docker-compose \
	&& chmod +x /usr/bin/docker-compose
