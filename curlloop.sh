#!/usr/bin/env bash

export APP="$1"
echo "curling URL $APP in a loop..."

while true
do 
    curl $APP 
    sleep 120
done
