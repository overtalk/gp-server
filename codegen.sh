#!/bin/bash

protoc -I ./ --go_out=plugins=grpc:protocol/ proto/*.proto
