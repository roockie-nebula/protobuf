#!/bin/sh

protoc \
  --go_out=. \
  --proto_path=.. \
  nullfloat.proto

mv ./github.com/roockie-nebula/protobuf/types/nullfloat/* .
rm -rf ./github.com
