## Makefile 使用指南

~~~~
    make是软件构建工具，主流的开源项目的主要构建手段。make构建工具主要是通过”Makefile“或”makefile“文件的形式来实现工程的自动化构建。通过target来检测文件之间的依赖关系，主要通过对比文件的修改时间来实现。我们使用make命令，通过Makefile将工程中的模块关联在一起，编译工程的源代码，然后把结果代码链接起来生成可执行文件或者库文件。
~~~~

>GNU make对make的标准功能（通过clean-room工程）进行了重新改写，并加入作者自认为值得加入的新功能，常和GNU编译系统一起被使用，是大多数GNU Linux默认安装的工具。

### 一、Makefile基本规则
```shell
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
