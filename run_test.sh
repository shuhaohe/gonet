#!/usr/bin/env bash

go test -race -c -v $1
export name=`basename $1`
sudo setcap CAP_NET_RAW=epi ./${name}.test
./${name}.test -test.v
#rm ${name}.test
