#!/usr/bin/env bash

if [ $USER == "root" ]; then
  /go/bin/staticcheck ./...
else
  $HOME/go/bin/staticcheck ./...
fi
