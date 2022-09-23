#!/bin/bash

PID=$(lsof -i:8089 | grep main | awk '{print $2}')
kill -9 ${PID}
nohup go run *.go 1>service.log 2>&1 &