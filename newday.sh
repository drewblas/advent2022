#!/bin/bash

# Error if no argument is given
if [ $# -eq 0 ]; then
    echo "Usage: newday.sh <day>"
    exit 1
fi

cp -a template day$1

mv day$1/day.go day$1/day$1.go

# Replace all instances of "dayXX" with "day$1"in the new directory
find day$1 -type f -exec sed -i "s/template/day$1/g" {} \;

