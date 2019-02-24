# 变量赋值

# 使用.PHONY定义Makefile的执行结构
# 使用.DEFAULT_GOAL定义默认执行target

.PHONY: clean one two three four five six seven eight

.DEFAULT_GOAL:= one

# 定义变量部分
nullName=
myName = "李雷"

one:clean two three
	@echo "start make file."

three: four five
	@echo "begin to use the variavle."
	@echo "nullName = "${nullName}
	@echo "myName = "$(myName)
	@echo $$


five: six seven
	


