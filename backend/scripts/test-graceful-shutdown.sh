#!/usr/bin/env bash

cd `dirname -- $0`; cd ..
echo "building..."
go build .
DEV_ROUTES=true IN_MEMORY_DB=true ./backend &
srvpid=$!
sleep 1

(
  curl localhost:3000/dev/delay/5
) &

sleep 1
echo 'sending SIGTERM to service...'
kill $srvpid
wait
