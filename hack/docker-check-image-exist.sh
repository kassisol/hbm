#!/bin/sh

IMAGE_REPO=$1
IMAGE_NAME=$2
TAG_VERSION=$3

if [ -z ${IMAGE_REPO} ]; then
	echo "Variable \"IMAGE_REPO\" not set"
	exit 1
fi
if [ -z ${IMAGE_NAME} ]; then
	echo "Variable \"IMAGE_NAME\" not set"
	exit 1
fi
if [ -z ${TAG_VERSION} ]; then
	echo "Variable \"TAG_VERSION\" not set"
	exit 1
fi

tags=`docker images | awk "/${IMAGE_REPO}\/${IMAGE_NAME}/ && match(\\$1, /^${IMAGE_REPO}\/${IMAGE_NAME}$/) { print \\$2 }"`

for tag in ${tags}; do
	if [ "$tag" == "${TAG_VERSION}" ]; then
		echo "Image ${IMAGE_REPO}/${IMAGE_NAME} with tag ${TAG_VERSION} already exist!"
		exit 1
	fi
done
