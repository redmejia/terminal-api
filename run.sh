#!/bin/bash

# clear the term
clear

echo "terminal"

# datatase enviroment variables
DBENV=.dbenv

# check if file exist
if [ -e "$DBENV" ]; then
	echo "$DBENV setting database enviroments"
	. .dbenv
	echo "run server"
	go run *.*go
else
	echo "$DBENV ...NOT FOUND"
	echo "...creating"
	touch .dbenv
	echo "file was creates edit with enviroment vars"
fi
