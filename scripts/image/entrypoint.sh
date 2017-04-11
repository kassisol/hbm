#!/bin/bash

HBM_LIB="/var/lib/hbm"
HBM="/usr/local/sbin/hbm"

if [ ! -f "${HBM_LIB}/data.db" ]; then
	${HBM} init
fi

exec ${HBM} server
