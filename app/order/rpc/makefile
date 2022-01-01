define checkdir
	@echo "创建目录$(1)"
    mkdir -p $(1)
endef

style=gozero

# rpc
rpc_name=order
rpc_dir=.
out_put=bin/server
# 生成proto模板
rpc-template:
	goctl rpc template -o $(rpc_dir)/$(rpc_name).proto

# 根据已有rpc文件生成服务
rpc-proto:
	goctl rpc proto -src $(rpc_dir)/$(rpc_name).proto  -dir $(rpc_dir) -style $(style)

build:
	go build -o $(out_put) -ldflags "-s -w" ./
	chmod u+x $(out_put)
clean:
	# rm $(out_put) *.bak *.~ -f