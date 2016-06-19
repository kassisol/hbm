#!/bin/sh

GIT_BRANCH=`git symbolic-ref --short -q HEAD 2> /dev/null`

echo ${GIT_BRANCH}
