.PHONY: proto build clean

PROTO_DIR := proto
RPC_OUT   := ./rpc

# 生成 proto 对应的 go / grpc 代码
# 生成结果: rpc/system/{common,user,role,dept,menu}.pb.go
#           rpc/system/{user,role,dept,menu}_grpc.pb.go
proto:
	protoc -I $(PROTO_DIR) \
		--go_out=$(RPC_OUT) \
		--go-grpc_out=$(RPC_OUT) \
		$(PROTO_DIR)/system/*.proto

# 编译整个项目
build:
	go build ./...

# 清理生成文件（谨慎使用）
clean:
	find $(RPC_OUT) -name '*.pb.go' -delete
