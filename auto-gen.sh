#!/bin/bash

MOD=`go mod why | awk 'END {print}'`
# protoc --go_out=plugins=grpc:. \
#     --gokit-micro_out=. \
#     --gokit-micro_opt=paths=source_relative,gentype=client,mod=$MOD \
#     --openapiv2_out=./swagger \
#     --openapiv2_opt=logtostderr=true \
#     -I./pb \
#      ./pb/version.proto 

protoc --go_out=plugins=grpc:. \
    --gokit-micro_out=. \
    --gokit-micro_opt=paths=source_relative,mod=$MOD \
    --openapiv2_out=./swagger \
    --openapiv2_opt=logtostderr=true \
    -I./pb \
    ./pb/version.proto 


