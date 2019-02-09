#!/usr/bin/env bash

for file in scripts/*
do
    if [[ -f $file ]]; then
        bash $file
        echo "Ran script: $file"
    fi
done