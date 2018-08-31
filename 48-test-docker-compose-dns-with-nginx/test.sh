#!/usr/bin/env bash -x

docker-compose up -d --scale api=3

for i in {1..10}; do
    curl localhost:8080
done

docker-compose down
