#!/bin/bash
#sudo docker run -i -t --name esseryu hqbfs:latest
var=`docker images hqbfs`

if [[ $var == *"hqbfs"* ]] && [[ $var == *"latest"* ]]
then
	echo "Hqb File Sever container exists for docker"
else
	echo "Build Hqb File Sever container for docker"
	tmp1=`docker build --tag hqbfs .`
fi

var1=`docker images mongo`

if [[ $var1 == *"mongo"* ]] && [[ $var1 == *"latest"* ]]
then
	echo "mongodb container exists for docker"
else
	echo "Get mongodb container for docker"
	tmp2=`docker pull mongo`
fi

docker stop hqbfs_d
docker stop hqbfs_s
docker rm hqbfs_d
docker rm hqbfs_s
cd ..
docker run --name hqbfs_d -d -v "hqbfs_db:/data/db" mongo:latest
docker run -it --name accesswall -d --volume=$(pwd):/go/src/github.com/jintech2ng/accesswall -p 8000:8000 --link hqbfs_d:mongo hqbfs:latest ./hqbfs_start.sh
