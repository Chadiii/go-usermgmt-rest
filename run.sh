#!/bin/bash

docker build -f Dockerfile -t usermgmt/rest .
docker run --rm -t -p 8081:8080 --name rest \
 -v $(pwd)/:/go/src/github.com/Chadiii/go-usermgmt-rest \
 usermgmt/rest