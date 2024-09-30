#!/bin/bash

# count fils in directory
count=$(find . -type f -o -type d | wc -l)

# *5
result=$((count * 5))

# print entr
printf "\t\tTotal files * 5: %d\v\n" "$result"