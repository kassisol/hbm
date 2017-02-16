#!/bin/bash

RPMBUILD_PATH="/srv/rpmbuild"
VOL_MOUNT_PATH="/tmp/hbm"

if [ ! -d ${VOL_MOUNT_PATH} ]; then
	echo "No volume mounted to path ${VOL_MOUNT_PATH}"

	exit 1
fi

cp  ${RPMBUILD_PATH}/RPMS/x86_64/*.rpm /tmp/hbm/
