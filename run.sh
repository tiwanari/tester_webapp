#!/bin/bash
set -xeu


docker build -t tiwanari/tester .
docker run --rm -p 5000:5000 tiwanari/tester
