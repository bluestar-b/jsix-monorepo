#!/bin/bash


directories=(
    "homepage/wserve"
    "Uploads"
    "sp2/web"
    "stats-api-rewrite"
)


for dir in "${directories[@]}"; do
  (cd "$dir" && go build)
  echo "Built in directory: $dir"
done

