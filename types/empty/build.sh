#!/bin/sh

protoc \
  --go_out=. \
  --proto_path=.. \
  empty.proto

mv ./github.com/roockie-nebula/protobuf/types/empty/* .
rm -rf ./github.com
