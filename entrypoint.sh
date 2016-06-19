#!/bin/sh

function copy_binary() {
	cp /go/bin/hbm /tmp/hbm/
}

function copy_rpm() {
	cp /srv/rpmbuild/RPMS/x86_64/*.rpm /tmp/hbm/
}


param=$1

if [ ${param} == "all" ]; then
	copy_binary
	copy_rpm
elif [ ${param} == "binary" ]; then
	copy_binary
elif [ ${param} == "rpm" ]; then
	copy_rpm
else
	echo "Parameter \"${param}\" not recognized"
fi
