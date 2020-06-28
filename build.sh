#!/usr/bin/env bash
GOOS=linux go build -o k8-web-terminal
#GOOS=darwin go build -o bin/k8-web-terminal-mac
#GOOS=windows go build -o bin/k8-web-terminal.exe

docker build -t  k8-web-terminal:test .
docker tag k8-web-terminal:test quay.xiaodiankeji.net/yangchun/k8-web-terminal:test
#docker push quay.xiaodiankeji.net/yangchun/k8-web-terminal:test
