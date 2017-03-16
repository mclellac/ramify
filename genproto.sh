#!/bin/sh

for f in ./services/**/*.proto; do
    echo "Generating protocol buffer for ${f}"
    protoc -I ./services ${f} --go_out=plugins=grpc:protos
done
