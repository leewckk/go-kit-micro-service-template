#!/bin/bash


if [ ! "$1" ] ; then
    echo 'usage :'
    echo './init.sh $modname'
    exit
fi

DEST=$1
go mod init $DEST 

#修改module的名称
SRC="github.com/leewckk/go-kit-micro-service-template"
# DEST=`go mod why | awk 'END {print}'`
go mod tidy 

sed -i "s#$SRC#$DEST#g" `grep "$SRC" -rl --include="*.go" .`
# sed -i "s#$SRC#$DEST#g" `grep "$SRC" -rl ./Makefile`
# sed -i "s#$SRC#$DEST#g" `grep "$SRC" -rl ./Dockerfile`

make 

