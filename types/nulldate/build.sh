#!/bin/sh

protoc \
  --go_out=. \
  --proto_path=.. \
  nulldate.proto

mv ./github.com/roockie-nebula/protobuf/types/nulldate/* .
rm -rf ./github.com
