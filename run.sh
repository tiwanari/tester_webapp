#!/bin/bash
set -xeu


if [[ "$(docker images -q tiwanari/tester:latest 2> /dev/null)" == "" ]]; then
    docker build -t tiwanari/tester .
fi

docker run --rm -p 5000:5000 tiwanari/tester
