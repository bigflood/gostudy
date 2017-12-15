#!/bin/bash

docker run -p 8080:8080 -e REDIS_ENDPOINT=192.168.99.100:6379 webchat-svr
