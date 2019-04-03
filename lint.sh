
#!/bin/bash

protoc -I . --lint_out=. proto/*.proto
