#!/bin/bash
docker kill backend
docker rm mackend
docker run --name backend -p 80:80 -v `pwd`/default.conf:/etc/nginx/conf.d/default.conf -d nginx:stable-alpine