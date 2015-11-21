#!/usr/bin/env bash

set -e

here="$(dirname "$BASH_SOURCE")"
cd $here
source env.sh
cd _vendor
mkdir -p src/github.com/cloudimmunity
ln -sf $BDIR src/github.com/cloudimmunity/system
#go get -d github.com/cloudimmunity/system
