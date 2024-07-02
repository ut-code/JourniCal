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

start vet
go vet ./...
assert_no_change vet

start fmt
go fmt ./...
assert_no_change fmt

start staticcheck
go run honnef.co/go/tools/cmd/staticcheck@latest ./...
if [ $? == 0 ]; then
  echo [Go tests] staticcheck passed.
else
  echo [Go tests] staticcheck failed.
  exit_code=1
fi

start test
go test -v ./pkg/...
go test -v ./app/...
assert_no_change test

start build
go build -n 2&>/dev/null
assert_no_change build

exit $exit_code
