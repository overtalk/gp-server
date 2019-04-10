#! /bin/bash

rm -rf ./model/xorm/

xorm reverse mysql root:12345678@tcp\(127.0.0.1:3306\)/gp_oj\?charset=utf8 $GOPATH/src/github.com/go-xorm/cmd/xorm/templates/goxorm  ./model/model

cd ./model/model/

rm `ls | grep -v "go" `

cd ..

mv model xorm