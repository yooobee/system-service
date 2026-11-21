### 1. go-zero生成命令
```md
yooome@me user-service % goctl rpc protoc rpc/system/system.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=. --style=goZero

yooome@anonymous user-service % goctl rpc protoc \
--go_out=./rpc \
--go-grpc_out=./rpc \
--zrpc_out=. \
--style=goZero \
rpc/system/system.proto

```
### 2. 生成sql命令
```md
goctl model mysql ddl -src ./model/UserInfo.sql -dir ./model --style=goZero
```
### 3. 多个service
```text
goctl rpc protoc rpc/system/system.proto \
  --go_out=./rpc \
  --go-grpc_out=./rpc \
  --zrpc_out=. \
  --style=goZero \
  -m

```
### 4. 单个role.proto文件
```bash
protoc -I . --go_out=./rpc rpc/system/role.proto
```
### 5. 最重要的命令
```bash
# 项目根目录执行
goctl rpc protoc proto/system/config.proto \
  --go_out=./rpc \
  --go-grpc_out=./rpc \
  --zrpc_out=. \
  --style=goZero \
  -m
```