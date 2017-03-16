#!/bin/sh

for f in ./protos/**/*.proto; do
    echo "Generating protocol buffer for ${f}"
    protoc -I ./protos ${f} --go_out=plugins=grpc:protos
done
