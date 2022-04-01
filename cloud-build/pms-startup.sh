#! /bin/bash

# start PMS
/init &

# wait for PMS to start
sleep 5

echo "starting go tests..."

# run go tests
go test