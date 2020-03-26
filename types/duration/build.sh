#!/bin/sh

protoc \
  --go_out=. \
  --proto_path=.. \
  duration.proto

mv ./github.com/roockie-nebula/protobuf/types/duration/* .
rm -rf ./github.com
