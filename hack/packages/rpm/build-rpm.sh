#!/bin/bash

RPMBUILD_PATH="/srv/rpmbuild"

mkdir -p ${RPMBUILD_PATH}

mkdir ${RPMBUILD_PATH}/BUILD
mkdir ${RPMBUILD_PATH}/RPMS
mkdir ${RPMBUILD_PATH}/SOURCES
mkdir ${RPMBUILD_PATH}/SPECS
mkdir ${RPMBUILD_PATH}/SRPMS
mkdir ${RPMBUILD_PATH}/tmp

echo "%_topdir  ${RPMBUILD_PATH}" > /root/.rpmmacros
echo "%_tmppath  ${RPMBUILD_PATH}/tmp" >> /root/.rpmmacros

cd  ${RPMBUILD_PATH}/SPECS

cp /usr/local/src/hbm.spec  ${RPMBUILD_PATH}/SPECS/

tar xvzf /tmp/hbm.tar.gz -C /usr/local/src/hbm

tar cvzf /srv/rpmbuild/SOURCES/hbm.tar.gz -C /usr/local/src hbm

VERSION=`/usr/local/src/hbm/hbm version | awk '{ print $2 }'`
if [ -z $VERSION ]; then
	echo "VERSION var is not set"

	exit 1
fi

rpmbuild -ba \
	--define '_release 1' \
	--define "_version ${VERSION}" \
	--define '_unitdir etc/systemd/system' \
	hbm.spec

#rpmlint hbm.spec ../SRPMS/hbm* ../RPMS/*/hbm*
