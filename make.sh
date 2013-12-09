#!/bin/sh

export GOPATH=`pwd`
go install blog
go install web
go install test
