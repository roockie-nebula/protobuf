#!/bin/sh

protoc \
  --go_out=. \
  --proto_path=.. \
  nullfloat.proto

mv ./github.com/dkfbasel/protobuf/types/nullfloat/* .
rm -rf ./github.com
