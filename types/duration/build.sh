#!/bin/sh

protoc \
  --go_out=. \
  --proto_path=.. \
  duration.proto

mv ./github.com/dkfbasel/protobuf/types/duration/* .
rm -rf ./github.com
