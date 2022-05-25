#! /bin/bash

# start PMS in the background
# TODO: use 'expect' to automate test start up
/init &

# wait for PMS to start
sleep 15
echo "starting go tests..."

# run go tests
cd /app/go-plex-client
go test