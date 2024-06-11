#!/usr/bin/env bash

# src of this function: https://developers.wano.co.jp/1883/ (modified)
function assert_no_change () {
  git diff --exit-code --quiet;
  if [ $? -eq 0 ];then
    echo Go tests: $1 passed;
  else
    echo Go tests: $1 failed;
    exit 1;
  fi
}
        
echo Starting go vet...
go vet ./...
assert_no_change vet
        
echo Starting go fmt...
go fmt ./...
assert_no_change fmt
        
echo Starting staticcheck...
go run honnef.co/go/tools/cmd/staticcheck@latest ./...
assert_no_change staticcheck
        
echo Starting go test...
go test -v ./...
assert_no_change test
        
echo Starting go build...
go build -n 2&>/dev/null
