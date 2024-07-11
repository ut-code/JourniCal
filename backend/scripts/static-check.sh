#!/usr/bin/env bash
cd $(dirname -- $0)
cd ..

exit_code=0
function assert_no_change () {
  git diff --exit-code --quiet;
  if [ $? -eq 0 ]; then
    echo [Go tests] $1 passed
  elif [ exit_code == 0 ]; then
    echo [Go tests] failed at: $1
    exit_code=1
  fi
}
function start () {
  echo [Go tests] starting $1...
}
function assert_zero () {
  if [ $1 == 0 ]; then
    echo [Go tests] $2 passed
  else
    echo [Go tests] failed at: $2
    exit_code=1
  fi
}

start vet
go vet ./...
assert_no_change vet

start fmt
go fmt ./...
assert_no_change fmt

start staticcheck
go run honnef.co/go/tools/cmd/staticcheck@latest ./...
assert_zero $? staticcheck

start test
go test ./pkg/...
assert_zero $? test
go test ./app/...
assert_zero $? test

start build
go build -n 2>/dev/null
assert_zero $? build

exit $exit_code
