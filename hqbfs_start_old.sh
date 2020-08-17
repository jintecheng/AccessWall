#!/bin/sh
sleep 15s
./wait-for-it.sh $MONGO_PORT_27017_TCP_ADDR:$MONGO_PORT_27017_TCP_PORT -- echo "mongod is up"
./hqbfs $MONGO_PORT_27017_TCP_ADDR $MONGO_PORT_27017_TCP_PORT
