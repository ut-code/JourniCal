#!/usr/bin/env bash

cd `dirname -- $0`
cd ..

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

start fmt
go fmt ./...
assert_no_change fmt
