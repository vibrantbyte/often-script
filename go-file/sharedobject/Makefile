# 定义变量 project
project = lib-vibrant.so

.PHONY: all vibrant

# 默认使用第一个真target，使用.DEFAULT_GOAL改变默认值
.DEFAULT_GOAL:= vibrant

all: vibrant
	@echo "all is execute."

vibrant:
	go build --buildmode=c-shared -o $(project) ./libvibrant.go