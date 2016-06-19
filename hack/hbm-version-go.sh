#!/bin/sh

VERSION=`hack/git-version.sh`

if [ ! -d "version" ]; then
	mkdir version
fi

if [ -f "version/version.go" ]; then
	rm -f version/version.go
fi

echo -e "package version\n\nvar VERSION = \"${VERSION}\"" > version/version.go
