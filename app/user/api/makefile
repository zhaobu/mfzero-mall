define checkdir
	@echo "创建目录$(1)"
    mkdir -p $(1)
endef

style=gozero

# api
api_name=user
api_dir=proto
out_put=bin/server
# 生成api文件
api-o:
	$(call checkdir,$(api_dir))
	goctl api -o $(api_dir)/$(api_name).api

# 根据已有api文件生成服务
api-go:
	goctl api go -api $(api_dir)/$(api_name).api -dir ./ -style $(style)

build:
	go build -o $(out_put) -ldflags "-s -w" ./
	chmod u+x $(out_put)
clean:
	rm $(out_put) *.bak *.~ -f