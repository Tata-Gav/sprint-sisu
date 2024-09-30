#!/bin/bash
if [-z "$1" ]; then
    echo "Error: No arguments supplied."
    exit 1
fi

count=$1

if [ "$count" -gt 100 ]; then
 count=100
fi

for ((i = 1; i<= count; i++))
do
 echo "$i times I've printed tatianagavrilova2"
done