## Makefile 使用指南

~~~~
    make是软件构建工具，主流的开源项目的主要构建手段。make构建工具主要是通过”Makefile“或”makefile“文件的形式来实现工程的自动化构建。通过target来检测文件之间的依赖关系，主要通过对比文件的修改时间来实现。我们使用make命令，通过Makefile将工程中的模块关联在一起，编译工程的源代码，然后把结果代码链接起来生成可执行文件或者库文件。
~~~~

>GNU make对make的标准功能（通过clean-room工程）进行了重新改写，并加入作者自认为值得加入的新功能，常和GNU编译系统一起被使用，是大多数GNU Linux默认安装的工具。

## Makefile简介

### 一、Makefile基本规则
```bash
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
```bash
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
```bash
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
```bash
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
```c
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

