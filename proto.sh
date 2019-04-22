#! /bin/bash

if [ $1 == "lint" ]
then
    protoc -I . --lint_out=. proto/*.proto
    exit 1
fi

if [ $1 == "doc" ]
then
    protoc -I . --doc_out=docs/ --doc_opt=markdown,introduction.md proto/*.proto
    exit 1
fi

if [ $1 == "code" ]
then
    protoc -I ./ --go_out=plugins=grpc:./ proto/*.proto
    exit 1
fi

echo "非法参数[ $1 ]" 
echo "  args : 
        - lint : 语法检查
        - doc  : 生成protobuf文档
        - code : 编译proto文件"