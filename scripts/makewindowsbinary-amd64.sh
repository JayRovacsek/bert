#!/usr/bin/env bash

env GOOS=windows GOARCH=amd64 go build -o binaries/comparison-amd64.exe comparison.go