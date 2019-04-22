#! /bin/bash

export CONFIG_FILE_PATH="$GOPATH/src/github.com/qinhan-shu/gp-server/config/" 
JUDGE_FILE_PATH="$GOPATH/src/github.com/qinhan-shu/gp-server/docker_data/judge/judgefile/" 

if [ $1 == "web" ]
then
    go run server/web/server.go 
    exit 1
fi

if [ $1 == "file" ]
then
    go run server/file/server.go -uploadPath $JUDGE_FILE_PATH
    exit 1
fi

echo "非法参数[ $1 ]" 
echo "  args : 
        - web  : 启动web服务器
        - file : 文件服务器"