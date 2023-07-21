#!/bin/sh
set +e
cd $1
make proto
rm -rf go.mod
rm -rf go.sum
go mod init $2
go mod tidy
#依赖注入编译
all_wires=$(find $1/src/services -name "wire.go" -type f)
for w in $all_wires; do
    #   echo "[INFO] ==> compile path:"$api_path
    wire $w
done
