# GG-IceCreamShop

GraphQL + gRPC Ice Cream Shop

## Getting Started

This project requires the `google.protobuf` proto package and expects to import from the `/usr/local/include` directory.

### Setting up the `google.protobuf` package & `protoc` binary

Download the applicable release on <a href="https://github.com/protocolbuffers/protobuf/releases" target="_blank">github</a>, and the downloaded zip archive contains the `bin` and `include` directory.

Extract the archive and copy the `bin/protoc` to your `GOBIN` path and copy the contents of the `include` directory to your `/usr/local/include` directory.