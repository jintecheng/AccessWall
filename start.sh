#!/bin/sh
sleep 15s
./wait-for-it.sh localhost:27017 -- echo "mongod is up"
./accesswall localhost 27017
