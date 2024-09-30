#!/bin/bash

# count fils in directory
count=$(find . | wc -l); printf "\t\vTotal files * 5: %d\v\n" "$((count * 5))
