#!/bin/sh
# 由于国内的原因，导致golang.org 的文件下载不下来，所有用这个脚本替代

# 第一部分 安装git
apk add git 

# 第二部分
rootdir=`pwd`
echo "this pwd is $rootdir"

golangpath="$rootdir/vendor/golang.org/x"
echo "golangpath :"$golangpath

if [ ! -d $golangpath ]; then
    
    mkdir -p  $golangpath
    cd $golangpath

    git clone https://github.com/golang/crypto.git
    git clone https://github.com/golang/sys.git

fi

echo "start install govendor..."
go get -u github.com/kardianos/govendor
echo "end install govendor..."

echo "start govendor sync -v..."
govendor sync -v
echo "end govendor sync -v..."

echo "start go build ..."
go build -o cosme .
ls -la cosme 
echo "end go build ..."

echo "completed !!!"