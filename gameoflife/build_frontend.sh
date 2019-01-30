#!/bin/bash

GOOS=js GOARCH=wasm go build -o ./static/main.wasm ./frontend
