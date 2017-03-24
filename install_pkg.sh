#!/usr/bin/env bash

if [ ! -f install_pkg.sh ]; then

echo 'install must be run within its container folder' 1>&2
exit 1

fi
CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

#gofmt -w src

## yaml
go get -u github.com/wendal/goyaml2
## json
go get -u github.com/bitly/go-simplejson

# setting fasthttp
go get -u github.com/valyala/fasthttp
# setting router
go get -u github.com/buaazp/fasthttprouter

export GOPATH="$OLDGOPATH"
echo 'finished'
