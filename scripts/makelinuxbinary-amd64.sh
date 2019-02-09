#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64 go build -o binaries/comparison-amd64 comparison.go