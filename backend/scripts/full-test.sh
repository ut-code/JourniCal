#!/usr/bin/env bash
cd $(dirname -- $0)
cd ..

go test ./test/...
