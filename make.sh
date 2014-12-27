#!/bin/sh

# compile program
export GOPATH=`pwd`
go install blog
go install exe/web
go install exe/babegarden

# generate blog txt
make -C data/blog/md
