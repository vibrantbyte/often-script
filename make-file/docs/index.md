## Makefile 使用指南

~~~~
    make是软件构建工具，主流的开源项目的主要构建手段。make构建工具主要是通过”Makefile“或”makefile“文件的形式来实现工程的自动化构建。通过target来检测文件之间的依赖关系，主要通过对比文件的修改时间来实现。我们使用make命令，通过Makefile将工程中的模块关联在一起，编译工程的源代码，然后把结果代码链接起来生成可执行文件或者库文件。
~~~~

>GNU make对make的标准功能（通过clean-room工程）进行了重新改写，并加入作者自认为值得加入的新功能，常和GNU编译系统一起被使用，是大多数GNU Linux默认安装的工具。

## Makefile简介

### 一、Makefile基本规则
```Makefile
target ...: prerequirements ...
    command ...
    ...
```
* target : 目标文件,可以有多个，可以是 .o 文件或者是可执行问价，甚至可以是一个标签。
* Prerequisites : 先决条件,可以是文件，也可以是另一个 target。
* command : 执行命令，linux shell 命令，当前用户权限下 u=x 的文件。
~~~
target + Prerequisites 可以组成一个递归的 ++『依赖关系』++ ，target 的先决条件定义在 prerequisites 中，而其生成规则又是由 command 决定的。如果包含多个规则的话，那么第一条规则就是整个 Makefile 的默认规则。
~~~

### 二、make的工作流程

1. 当执行make命令时，查找当前目录的Makefile文件。（分为加target和不加target）
2. 如果不加target将执行当前目录下的Makefile文件的第一个target。
3. 如果加了target将执行输入的target，如果找不到 "make: *** No rule to make target `prf'.  Stop."。
4. target正确的情况下，如果目标不存在，则寻找对应的 .o 文件。
5. 如果 .o 文件不存在，则寻找 .o 的依赖关系以生成它。

~~~
Makefile有默认生成规则，但我们常会自己编写规则，目的是方便自定义、便于移植、便于交叉编译、便于调试。
~~~

### 三、Makefile中的变量
```Makefile
# 注意变量的值是允许空格的
name = value
#使用变量
$(name)
```
value值过长或者command过长可以使用”\“来作为换行符，将内容连在一起，其实不是真正的换行，只是一种显示换行，执行的时候不会进行换行。

## Makefile综述
~~~~
文件指示：在一个 Makefile 里面可以制定另一个 makefile，类似于 C 的 include
Makefile 还可以做条件包含动作，类似于 #if。Makefile 可以定一个变量为一个多行的命令。
Makefile 里面只有行注释而没有段注释。注释采用 # 开头。如果要使用 # 字符，则需要转义，写成 “\#”。
Makefile 规则内容里所有的 shell 命令都要以制表符 Tab 开头，注意，空格符是不行的。
~~~~
> 默认的 make 文件名为：GNUmakefile, makefile, Makefile，当敲入 make 命令时，会自动搜寻这几个文件。约定俗成使用最后一个。

### 引用其他Makefile
基本语法：
```Makefile
include filename ...        # 不允许 include 失败
-include filename ...       # 允许 include 失败
```
> 可以包含路径或者通配符，一行可以包含多个文件。

如果未指定绝对路径或者相对路径，那么 make 会按照一下的顺序去寻找：
1. 当前目录
2. 制定 make 时，在 -I 或者 --include-dir 的参数下寻找
3. \<prefix>/include（一般是 /usr/local/bin 或 /usr/include）

> 建议还是手动指定吧，自动搜寻意外可能太多了。

### 环境变量 MAKEFILES
> 这里主要是要提醒：不要设置这个环境变量，否则会影响全局的 make include 动作。

### Make 的几个工作方法
> .DEFAULT_GOAL := all  指定默认target

**通配符**
> make 支持三个通配符：*, ?, [...]。可以用在规则中，也可以用在变量中。

**伪目标**
> 伪目标就是 Makefile 里面颇为常见的 .PHONY 标识，比如：".PHONY: clean"，表示这个规则名并不代表一个真实存在的、需要生成的文件名，而只是一条纯粹的规则。实例：[Makefile文件](../go-project/Makefile)

* 真目标的特点是：如果目标存在，才会被执行；不存在报错：
```bash
xiaoyueyadeMacBook-Pro:go-project xiaoyueya$ make testing
make: *** No rule to make target `testing'.  Stop.
```
* 伪目标的特点是：无视目标是否存在，必然执行.
```bash
xiaoyueyadeMacBook-Pro:go-project xiaoyueya$ make clean
make: Nothing to be done for `clean'.
```
除了 make clean 之外，伪目标还有另一种使用场景，就是一个 make 动作，实际上生成了多个目标。比如：
```Makefile
.PHONY: all
# 包含了生成目标文件，以及安装动作
all: exe install    
```

**总结：使用.PHONY的目的可以理解成先定义规则，可以先部分不实现来完成整个make过程。**

### 多目标
> 规则的冒号前面可以有多个 target，表示多个 target 共用这条规则。

### 自动生成依赖关系
~~~
如果我们使用中规中矩的 makefile 写法，那么对于每个源文件都要好好写头文件依赖关系，从而在头文件更新的时候，可以自动重新编译依赖于这个头文件的源文件。
这实在是太麻烦了。好在 gcc 里有一个 -MM（注意不是 “-M”） 的选项，可以分析出 .c 文件依赖的头文件并且打印出来。因此制作 Makefile 的时候，就可以利用这一特性自动生成依赖。
~~~

实现方法有很多，这里贴出我自己使用的例子，也可以参见我的工程代码：
```Makefile
EXCLUDE_C_SRCS =#
C_SRCS = $(filter-out $(EXCLUDE_C_SRCS), $(wildcard *.c))
C_OBJS = $(C_SRCS:.c=.o)

$(C_OBJS): $(C_OBJS:.o=.c)
    $(CC) -c $(CFLAGS) $*.c -o $*.o
    @$(CC) -MM $(CFLAGS) $*.c > $*.d  
    @mv -f $*.d $*.d.tmp  
    @sed -e 's|.*:|$*.o:|' < $*.d.tmp > $*.d  
    @sed -e 's/.*://' -e 's/\\$$//' < $*.d.tmp | fmt -1 | sed -e 's/^ *//' -e 's/$$/:/' >> $*.d
    @rm -f $*.d.tmp 
```

## 书写命令(command)
### 命令执行
| 符号 | 位置 | 含义 |
| :----| :---- | :---- |
| null | command首位 | 打印当前执行指令，失败打断 |
| @ | command首位 | 执行的时候，不打印这条命令语句，可以节省屏幕内容，减少无用信息输出 |
| - | command首位 | 无视这条命令的返回值是否为成功（0），永远都成功 |
| \ | command中间 | 将一行很长的命令展示成多行，在Makefile中形成假换行，主要目的是格式化，执行时其实是一行 |
| -n/--just-print | make 后紧跟 | 不是真正执行make，只是将过程打印出来 |

### 嵌套执行 make

&#8195;&#8195;在 Makefile 里可以到另一个目录下执行 make，执行方式类似于普通的命令调用，但特别的是，make 可以识别出这是一条嵌套 make 指令，从而在 shell 中打印出 “专项哪里哪里 make” 的提示语法为：

```Makefile
subsystem:
    $(MAKE) -C subdir
```

这个做法的主要好处是可以向下级 Makefile 传递变量或者语法：
```Makefile
export VARIABLE ... # 将相应变量变成当前 make 操作的全局变量
```

也恶意直接指定变量的值：
```Makefile
export VARIABLE = value
```
如果要传递所有变量（不推荐），直接写 export 就好。

> 注意由两个系统变量 SHELL 和 MAKEFLAGS 是永远传递的。  
 此外还有一个全局变量 MAKELEVEL 用来表示当前的嵌套层数。

### 定义命令包
命令包类似于宏、子函数等等。使用 define 来定义，以 endif 结束，比如：
```Makefile
define run-yacc
    yacc $(firstword $^)
    mv y.tab.c $@
endif
```
> 注意如果是命令的话，需要以制表符开头。调用这个命令包的方式为：$(run-yacc)

## 使用变量
### 变量赋值
Makefile中的变量可以只声明，不赋值，例如：（[variable使用](../variable-project/Makefile)）。主要使用规则如下:

| 符号 | 位置 | 含义 |
| :---- | :---- | :---- |
| $()/${}  | command中使用 | ()/{}内部为variable name，获取变量内容 |
| $$ | command中使用 | 获取$符号 |
| := | command之外 | 重新赋值变量，避免使用未定义的变量 |
| += | command之外 | “追加” 值。如果右侧有变量未定义，则等价于 “:=” |
| ?= | command之外 |  如果等号左侧的变量未定义，则使用等号右边内容定义，即：@(?=) |
| $@ | command中使用 |  表示当前规则的编译目标 |
| $^ | command中使用 |  表示当前规则的所有依赖文件 |
| $$< | command中使用 |  表示当前规则的第一个依赖 |

位置@(?=) 
```Makefile
ifeq($(some_var), undefined)
    some_var = some_val
endif
```

### 定义一个空格变量
```Makefile
NULL_STR :=#
SPACE_STR := $(NULL_STR) # end of line
```
> 注意第二行的注释与 “)” 之间是包含一个空格的。注释的 “#” 必须有，否则不会定义一个空格出来。

### 变量替换
> 第一个方式为：$(var: .o = .c)，意思是将等号左边的字符换成右边的字符  
第二个方式为所谓的 “静态模式”：$(var: %.o = %.c)

### 把变量值作为变量
> 这很类似于指针，只是地址值变成了变量值。可以用变量值生成变量名，比如：a := $($(var)) 或者是 $($(var)_$(idx)) 之类的写法。

### override
> 在命令行调用 make 时，可以直接指定某个变量的全局值，使得它在整个 make 的过程中一直不变。为了防止这个特性，可以使用这个关键字来处理：  
override \<variable> = \<value>  
等号也可以用 := 和 ?=

### 目标变量（局部变量）
如果某条约束里面不想使用已经定义了的全局变量，可以这样写：
```Makefile
prog: CFLAGS = -g
prog: a.o b.o
    $(CC) $(CFLAGS) a.o b.o
```

## 条件判断

### 语法
```Makefile
<条件语句>
<true 执行语句>
else
<false 执行语句>
endif
```
其中条件语句有四种情形：
> 表示是否相等

```Makefile
ifeq (<arg1>, <arg2>)    # 推荐
ifeq '<arg1>' '<arg2>'
ifeq "<arg1>" "<arg2>"
```
> 表示是否不等，上面的 ifeq 换成 ifneq

> ifdef

> ifndef


## 使用函数

### 常规函数
> 字符串替换
```Makefile
$(subst <from>, <to>, <text>)
```
> 模式字符串替换
```Makefile
$(patsubst <pattern>, <replacement>, <text>)
```
> 去开头和结尾的空格
```Makefile
$(strip <string>)
```
> 查找字符串
```Makefile
$(findstring <find>, <in>)
```
> 反过滤
```Makefile
$(filter-out <pattern_or_string>, <text>)
```
> 排序（单词升序）
```Makefile
$(sort <list>)
```
> 取单词
```Makefile
$(word <n>, <text>)
```
> 取单词串
```Makefile
$(wordlist <n_start>, <n_end>, <text>)
```
> 单词个数统计
```Makefile
$(words <text>)
```
> 去掉每个单词的最后文件名部分，只剩下目录部分
```Makefile
$(dir <names ...>)
```
> 去掉每个单词的目录部分，只剩下文件名部分
```Makefile
$(notdir <names ...>)
```
> 读取各文件名的后缀
```Makefile
$(suffix <names ...>)
```
> 加后缀
```Makefile
$(addsuffix <suffix>, <names ...>)
```
> 加前缀  
加前后缀在动态创建局部变量很有用
```Makefile
$(addprefix <prefix>, <names ...>)
```
> 连接字符串
```Makefile
$(join <list1>, <list2>)
```

### for 循环
~~~
$(foreach <var>, <list>, <text>)
这其实是一个函数，作用是：将 list 的单词逐一取出，放到 var 指定的变量中，然后执行 text 的表达式。返回值则是 text 的最终执行值。
~~~

### Shell 函数
~~~
执行 shell 命令，并且将 stdout 作为返回值返回，如：
contents := $(shell ls -la)
~~~

### 控制 make 输出
~~~
$(error <text ...>)
$(warning <text ...>)
这也同时是调试和定位 make 的好方法。
~~~

### 判断文件是否存在
```Makefile
ifeq ($(FILE), $(wildcard $(FILE)))
...
endif
```
