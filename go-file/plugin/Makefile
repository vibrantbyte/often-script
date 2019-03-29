# 定义变量 project  -buildmode=plugin 只支持linux，请注意
project = lib-usertree-plugin.so

.PHONY: all lib_plugin

# 默认使用第一个真target，使用.DEFAULT_GOAL改变默认值
.DEFAULT_GOAL:= lib_plugin

all: lib_plugin
	@echo "all is execute."

lib_plugin:
    go build --buildmode=plugin -o $(project) ./rrc_user_tree_plugin.go