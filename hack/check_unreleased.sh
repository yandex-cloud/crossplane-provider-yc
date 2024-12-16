#!/bin/bash

CHANGES_DIR=".changes/unreleased"

if [ ! -d "$CHANGES_DIR" ]; then
    echo "Error: Directory $CHANGES_DIR does not exist"
    exit 1
fi

files=$(ls -A "$CHANGES_DIR" | grep -v "^\.gitkeep$")

if [ ! -z "$files" ]; then
    echo "Error: Directory $CHANGES_DIR contains unexpected files:"
    echo "$files"
    exit 1
fi

echo "Check passed: $CHANGES_DIR contains only .gitkeep"
exit 0
