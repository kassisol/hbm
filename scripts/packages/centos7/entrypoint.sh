#!/bin/bash

VERSION=$1
RELEASE=$2

VERSION=`echo ${VERSION} | sed 's/-/_/'`

cd ${RPMBUILD_PATH}/SPECS

rpmbuild -ba \
	--define "_version ${VERSION}" \
	--define "_release ${RELEASE}" \
	--define '_unitdir etc/systemd/system' \
	hbm.spec

mkdir -p /tmp/dist
cp ${RPMBUILD_PATH}/RPMS/x86_64/*.rpm /tmp/dist/

#rpmlint hbm.spec ../SRPMS/hbm* ../RPMS/*/hbm*
