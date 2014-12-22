#!/bin/sh

# compile program
export GOPATH=`pwd`
go install blog
go install web
go install test
go install briabby

# generate blog txt
make -C data/blog/md
