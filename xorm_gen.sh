#! /bin/bash

rm -rf ./model/xorm/

xorm reverse mysql ${YOUR_DB_USERNAME}:${YOUR_DB_PWD}@tcp\(127.0.0.1:3306\)/${YOUR_DB_NAME}\?charset=utf8 $GOPATH/src/github.com/go-xorm/cmd/xorm/templates/goxorm  ./model/model

cd ./model/model/

rm `ls | grep -v "go" `

cd ..

mv model xorm