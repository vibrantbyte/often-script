# 此项目为go的hello world项目。使用make进行编译

# 定义变量 project
project = ”my name is go project“

.PHONY: all clean

# 默认使用第一个真target，使用.DEFAULT_GOAL改变默认值
.DEFAULT_GOAL:= testall

all: production
	@echo "all is execute."

production: 
	echo $(project)

testall: create production
	@echo "test all project."

create:
	@echo "create a file." 