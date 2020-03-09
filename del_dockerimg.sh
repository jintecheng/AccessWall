#!/bin/bash

val1=`docker images hqbfs`

if [[ $val1 == *"hqbfs"* ]] && [[ $val1 == *"latest"* ]]
then
	va=`docker rmi -f hqbfs:latest`
fi

val2=`docker images hqbfs_shell`

if [[ $val2 == *"hqbfs_shell"* ]] && [[ $val2 == *"latest"* ]]
then
	vaa=`docker rmi -f hqbfs_shell:latest`
fi
