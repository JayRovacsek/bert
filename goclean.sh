#!/usr/bin/env bash

go clean -i -n

for file in binaries/*
do
    if [[ -f $file ]]; then
        rm $file
        echo "rm -f $pwd/$file"
    fi
done