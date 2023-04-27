#!/bin/sh

set -e

make vendor
if ! git diff --exit-code go.mod go.sum; then
	echo "please run \`make mod\` and check in the changes"
	exit 1
fi
