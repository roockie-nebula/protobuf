#!/bin/sh

protoc \
  --go_out=. \
  --proto_path=.. \
  nullint.proto

mv ./github.com/roockie-nebula/protobuf/types/nullint/* .
rm -rf ./github.com
