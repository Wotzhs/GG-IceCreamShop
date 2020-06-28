#!/bin/sh

apk add make protobuf

wget https://github.com/protocolbuffers/protobuf/releases/download/v3.12.3/protoc-3.12.3-linux-x86_64.zip

mkdir protoc /usr/local/include

unzip protoc-3.12.3-linux-x86_64.zip -d protoc

mv protoc/include/* /usr/local/include