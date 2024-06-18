#!/usr/bin/env bash

cd $(dirname -- $0)
cd ..

go build -o ./dist/backend -ldflags="-s -w" -trimpath .
