#!/usr/bin/env bash

ROOTDIR=$(dirname $0)/../../..
cd $(dirname $0)

if [ -d "build" ]; then
	rm -rf build
fi
mkdir -p build

cp ${ROOTDIR}/contrib/init/systemd/hbm.service build/
cp ${ROOTDIR}/contrib/init/systemd/hbm.socket build/
cp ${ROOTDIR}/bin/hbm build/

go run ${ROOTDIR}/gen/man/genman.go
cp -r /tmp/hbm/man build/

go run ${ROOTDIR}/gen/shellcompletion/genshellcompletion.go
cp -r /tmp/hbm/shellcompletion build/
