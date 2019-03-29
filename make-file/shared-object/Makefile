# 此项目为一个没有文件的kong项目，完全使用网络来构建一个应用程序
# 定义变量 project名称
project=vibrant

.PHONY: all clean wget build install goal

# 默认使用第一个真target，使用.DEFAULT_GOAL改变默认值
.DEFAULT_GOAL:= all

all: before wget build install clean
	@echo "all is execute."

before:
	rm -rf ./lib${project}.so

wget: goal
	@echo "get from internet"
	@wget https://raw.githubusercontent.com/vibrantbyte/often-script/master/go-file/sharedobject/libvibrant.go
	@wget https://raw.githubusercontent.com/vibrantbyte/lua2go/master/lua/lua2go.lua

build: goal
	@echo "execute script"
	go build -ldflags "-s -w" -buildmode=c-shared -o lib${project}.so ./libvibrant.go
	# 增加执行权限
	chmod g+x,u+x,o+x libvibrant.so

install: goal
	@echo "install application"

clean: goal
	@echo "clean successful"
	rm -rf ./libvibrant.go ./libvibrant.h
