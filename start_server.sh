#!/bin/bash

docker build -t fibonacci .

docker run --restart=always -p 8080:8080 fibonacci