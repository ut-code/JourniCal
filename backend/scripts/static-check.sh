#!/usr/bin/env bash
cd $(dirname -- $0)
cd ..

rm test/*.db

./mutable.sh
go vet ./...
go run honnef.co/go/tools/cmd/staticcheck@latest ./...
go test ./pkg/...
ENV_FILE=test go test ./app/...
go build -n 2>/dev/null

rm test/*.db
