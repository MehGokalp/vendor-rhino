#!/bin/bash

if [ -n "$(gofmt -l .)" ]; then
    echo "There is unformatted code, you should use `go fmt ./\.\.\.` to format it."
    gofmt -d .
    exit 1
else
    echo "Codes are formatted."
    exit 0
fi
