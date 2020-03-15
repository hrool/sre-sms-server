#!/bin/bash

docker build --target builder -t  sre-sms-server:builder . &&  docker image prune -f
docker build -t sre-sms-server . &&  docker image prune -f

# docker push
