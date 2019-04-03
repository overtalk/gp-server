#!/bin/bash

protoc -I . --doc_out=docs/ --doc_opt=markdown,introduction.md proto/*.proto
