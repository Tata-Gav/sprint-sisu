#!/bin/bash
echo "find Everything that starts with an 'a': "
find . -type f -name 'a*'

echo "All the files ending with a 'z': "
find . -type f -name '*z'

echo "All files starting with 'z' and ending with 'a':"
find . -type f -name 'z*a'