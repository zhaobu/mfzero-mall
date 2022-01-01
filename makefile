define checkdir
	@echo "创建目录$(1)"
    mkdir -p $(1)
endef


clean:
	find . -name logs|xargs sudo rm -rf
	find . -name bin|xargs sudo rm -rf

style=gozero
service=demo
# api
api_name=$(service)
api_dir=api/$(api_name)

# 生成api文件
api-o:
	$(call checkdir,$(api_dir))
	goctl api -o $(api_dir)/$(api_name).api

# 根据已有api文件生成服务
api-go:
	goctl api go -api $(api_dir)/$(api_name).api -dir $(api_dir) -style $(style)



# rpc
rpc_name=$(service)
rpc_dir=app/$(rpc_name)/service

# 生成proto模板
rpc-template:
	goctl rpc template -o $(rpc_dir)/$(rpc_name).proto

# 根据已有rpc文件生成服务
rpc-proto:
	goctl rpc proto -src $(rpc_dir)/$(rpc_name).proto  -dir $(rpc_dir) -style $(style)


# model
table=user
sql_name=$(table).sql
model_name=$(table)
model_dir=app/$(service)/model
db_type=mysql
db_url=lw:12345@tcp(10.133.236.150:3306)/mf_demo

# 根据已有sql文件生成
model-ddl:
	$(call checkdir,$(model_dir))
	goctl rpc template -o $(rpc_dir)/$(rpc_name).proto
	goctl model mysql ddl -src $(model_dir)/$(sql_name) -dir $(model_dir) -c

# 根据db里面的表结构生成
model-url:
	$(call checkdir,$(model_dir))
	goctl model $(db_type) datasource -url="$(db_url)" -table="$(table)" -c -dir $(model_dir)

# template


