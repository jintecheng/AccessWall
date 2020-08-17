#!/bin/bash
#sudo docker run -i -t --name esseryu hqbfs:latest
docker stop hqbfs_d
docker stop hqbfs_s
docker rm hqbfs_d
docker rm hqbfs_s
