#!/bin/sh

CONTRACT_PATH=stats_article_main.proto
OUT_PATH=../src/api_stats_rpc/pb

mkdir -p ${OUT_PATH}
cd spec
protoc --go_out=${OUT_PATH} \
  --go_opt=paths=source_relative \
  --go-grpc_out=require_unimplemented_servers=false:${OUT_PATH} \
  --go-grpc_opt=paths=source_relative \
  ${CONTRACT_PATH}
cd ..
