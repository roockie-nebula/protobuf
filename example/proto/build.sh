#!/bin/sh

protoc \
  --go_out=plugins=grpc:../domain \
  --proto_path=. \
  --proto_path=../../../../ \
  *.proto

protoc-go-tags --dir=../domain
