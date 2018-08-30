#!/usr/bin/env bash

VERSION=$1
RELEASE=$2

VERSION=${VERSION//-/_}

cd "${RPMBUILD_PATH}/SPECS" || exit

rpmbuild -ba \
	--define "_version ${VERSION}" \
	--define "_release ${RELEASE}" \
	--define '_unitdir etc/systemd/system' \
	hbm.spec

mkdir -p /tmp/dist
cp ${RPMBUILD_PATH}/RPMS/x86_64/*.rpm /tmp/dist/

#rpmlint hbm.spec ../SRPMS/hbm* ../RPMS/*/hbm*
