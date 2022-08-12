#!/bin/bash

#修改module的名称
SRC="micro-service-sample"
DEST=`go mod why | awk 'END {print}'`

rm ./go.mod 
rm ./go.sum

go mod init $DEST 
go mod tidy 

sed -i "s/$SRC/$DEST/g" `grep "$SRC" -rl --include="*.go" .`
sed -i "s/$SRC/$DEST/g" `grep "$SRC" -rl --include="*.md" .`
sed -i "s/$SRC/$DEST/g" `grep "$SRC" -rl ./Makefile`
sed -i "s/$SRC/$DEST/g" `grep "$SRC" -rl ./Dockerfile`

make 

