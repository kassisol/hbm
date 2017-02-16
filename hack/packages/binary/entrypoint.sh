#!/bin/bash

VOL_MOUNT_PATH="/tmp/hbm"

if [ ! -d ${VOL_MOUNT_PATH} ]; then
	echo "No volume mounted to path ${VOL_MOUNT_PATH}"

	exit 1
fi

cp /usr/local/src/hbm-*-linux-amd64.tar.gz /tmp/hbm/
