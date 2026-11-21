### 1. 目录结构

```text
system-service
├── proto                  # 手写 proto 源文件（唯一维护入口）
│   └── system
│       ├── common.proto   # 通用分页 / Empty
│       ├── user.proto     # UserService   + 用户消息
│       ├── role.proto     # RoleService   + 角色消息
│       ├── dept.proto     # DeptService   + 部门消息
│       └── menu.proto     # MenuService   + 菜单消息
├── rpc
│   └── system             # protoc 生成的代码（不要手改）
│       ├── common.pb.go
│       ├── user.pb.go
│       ├── user_grpc.pb.go
│       ├── role.pb.go
│       ├── role_grpc.pb.go
│       ├── dept.pb.go
│       ├── dept_grpc.pb.go
│       ├── menu.pb.go
│       └── menu_grpc.pb.go
├── internal
│   ├── logic
│   │   ├── userservice    # package userservicelogic
│   │   ├── roleservice    # package roleservicelogic
│   │   ├── deptservice    # package deptservicelogic
│   │   └── menuservice    # package menuservicelogic
│   └── server
│       ├── userservice    # UserServiceServer
│       ├── roleservice    # RoleServiceServer
│       ├── deptservice    # DeptServiceServer
│       └── menuservice    # MenuServiceServer
├── main.go                # 一个 gRPC Server 注册四个服务
└── Makefile               # make proto / make build
```

四个 proto 文件都是 `package system; option go_package = "./system";`，
生成后在同一个 Go 包里，直接使用 `system.RegisterXxxServiceServer`。

### 2. 安装生成工具（首次）

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

# 确保 protoc 能找到插件
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --version
```

> 注意：`protoc-gen-go-grpc` 需要 Go >= 1.24；本地 Go 版本较低时请固定使用 v1.5.1。

### 3. 生成 proto 代码（推荐）

```bash
# 在项目根目录执行
make proto
```

等价于手动执行：

```bash
protoc -I proto \
  --go_out=./rpc \
  --go-grpc_out=./rpc \
  proto/system/common.proto \
  proto/system/user.proto \
  proto/system/role.proto \
  proto/system/dept.proto \
  proto/system/menu.proto
```

也可以使用通配符（依赖 shell 展开）：

```bash
protoc -I proto --go_out=./rpc --go-grpc_out=./rpc proto/system/*.proto
```

生成结果统一在 `rpc/system/` 下，所有文件同属 `system` 包。

### 4. 只生成某个 proto 的 pb.go（不生成 grpc）

```bash
protoc -I proto --go_out=./rpc proto/system/role.proto
```

### 5. go-zero 生成 model

```bash
goctl model mysql ddl -src ./model/UserInfo.sql -dir ./model --style=goZero
```

### 6. goctl 一次性生成 rpc 骨架（可选）

```bash
# 注意：多个 service 在同一 proto 文件时才需要 -m；
# 当前按领域拆分 proto，推荐直接使用上面的 protoc 命令，
# 逻辑/服务端目录已按 userservice、roleservice、deptservice、menuservice 组织。
goctl rpc protoc proto/system/user.proto \
  --go_out=./rpc \
  --go-grpc_out=./rpc \
  --zrpc_out=. \
  --style=goZero
```

### 7. 编译

```bash
make build
# 或
go build ./...
```

### 8. 启动

```bash
go run main.go -f etc/system.yaml
```

### 9. 新增一个业务领域（如 dict 字典）

1. 在 `proto/system/` 下新增 `dict.proto`，保持：

   ```proto
   syntax = "proto3";
   package system;
   option go_package = "./system";
   ```

2. 新增 `internal/logic/dictservice/` 与 `internal/server/dictservice/`。
3. 在 `main.go` 中追加：

   ```go
   system.RegisterDictServiceServer(grpcServer, dictserver.NewDictServiceServer(ctx))
   ```

4. 执行 `make proto`。
