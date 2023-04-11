#!/bin/bash

if [ -f ./netcat ]; then
    rm netcat
    echo "removing old exe"
fi

go build -race $1
./netcat

