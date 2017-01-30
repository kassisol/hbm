#!/bin/sh

HBM_SRC_PATH="/go/src/github.com/kasssol/hbm"

mkdir -p /srv/rpmbuild/

mkdir /srv/rpmbuild/BUILD
mkdir /srv/rpmbuild/RPMS
mkdir /srv/rpmbuild/SOURCES
mkdir /srv/rpmbuild/SPECS
mkdir /srv/rpmbuild/SRPMS
mkdir /srv/rpmbuild/tmp

echo '%_topdir /srv/rpmbuild' > /root/.rpmmacros
echo '%_tmppath /srv/rpmbuild/tmp' >> /root/.rpmmacros

cd /srv/rpmbuild/SPECS

mkdir /tmp/hbm
cp /go/bin/hbm /tmp/hbm/
cp ${HBM_SRC_PATH}/contrib/init/systemd/hbm.service /tmp/hbm/

tar cvzf /srv/rpmbuild/SOURCES/hbm.tar.gz -C /tmp hbm

cp ${HBM_SRC_PATH}/hack/hbm.spec /srv/rpmbuild/SPECS/

VERSION=`/tmp/hbm/hbm version | awk '{ print $2 }'`

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
