#!/bin/sh

VERSION_INI="0.1.0"

GIT_TAG_VERSION=`git describe --tags 2> /dev/null`

if [ -z ${GIT_TAG_VERSION} ]; then
	VERSION=${VERSION_INI}
else
	VERSION=`echo ${GIT_TAG_VERSION} | awk -F'-' '{ print $1 }'`
fi

echo ${VERSION}
