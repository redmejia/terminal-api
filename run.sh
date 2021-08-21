#!/bin/bash

echo "terminal"

DBENV=.dbenv

if [ -e "$DBENV" ]; then
	echo "$DBENV setting database enviroments"
	. .dbenv
	echo "run server"
	go run *.*go
else
	echo "...NOT FOUND"
	echo "create database enviroment variable file"
fi
