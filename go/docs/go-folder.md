## go 目录结构说明
~~~
    golang集多编程范式之大成者，使开发者能够快速的开发、测试、部署程序，支持全平台静态编译。go具有优秀的依赖管理，高效的运行效率，庞大的第三方库支持以及在国内持续的增长势头。  
    作为开发者的我们也将不得不重视这门语言的兴起。首先向大家讲解一下go语言开发环境的目录结构，让我们更清楚的认识它。
~~~
## 一、goroot开发包目录
~~~
    当我们安装好后，会在安装目录出现一个go/文件夹，如果是windows目录应在再C:/go下（默认），如果是unix/linux一般会在/usr/local/go下，这个目录是unix software resource的含义。
~~~
```bash
# liunx上目录位置
chao@chao-PC:/usr/local/go$ pwd
/usr/local/go

# 主要目录包含如下图，分别进行说明：
```
![go-folder](./images/go-folder.png)
### 1、api文件夹
~~~api
    存放Go API检查器的辅助文件。其中，go1.1.txt、go1.2.txt、go1.3.txt和go1.txt文件分别罗列了不同版本的Go语言的全部API特征；except.txt文件中罗列了一些（在不破坏兼容性的前提下）可能会消失的API特性；next.txt文件则列出了可能在下一个版本中添加的新API特性。
~~~
### 2、bin文件夹
~~~bin
    存放所有由官方提供的Go语言相关工具的可执行文件。默认情况下，该目录会包含go和gofmt这两个工具。
~~~
### 3、doc文件夹
~~~doc
    存放Go语言几乎全部的HTML格式的官方文档和说明，方便开发者在离线时查看。
~~~
### 4、misc文件夹
~~~misc
    存放各类编辑器或IDE（集成开发环境）软件的插件，辅助它们查看和编写Go代码。有经验的软件开发者定会在该文件夹中看到很多熟悉的工具。
~~~
查看：
```bash
chao@chao-PC:/usr/local/go/misc$ ls
android  benchcmp  chrome   git  linkcheck  sortac  tour
arm      cgo       editors  ios  nacl       swig    trace
```
### 5、pkg文件夹
~~~pkg
    用于在构建安装后，保存Go语言标准库的所有归档文件。pkg文件夹包含一个与Go安装平台相关的子目录，我们称之为“平台相关目录”。例如，在针对Linux 32bit操作系统的二进制安装包中，平台相关目录的名字就是linux_386；而在针对Windows 64bit操作系统的安装包中，平台相关目录的名字则为windows_amd64。
    Go源码文件对应于以“.a”为结尾的归档文件，它们就存储在pkg文件夹下的平台相关目录中。  
    值得一提的是，pkg文件夹下有一个名叫tool的子文件夹，该子文件夹下也有一个平台相关目录，其中存放了很多可执行文件。关于这些可执行文件的用途，读者可参见附属于本书的Go命令教程。
~~~
查看：
```bash
chao@chao-PC:/usr/local/go/pkg$ ls
include      linux_amd64_dynlink  linux_amd64_shared              tool
linux_amd64  linux_amd64_race     linux_amd64_testcshared_shared
```
### 6、src文件夹
~~~src
    存放所有标准库、Go语言工具，以及相关底层库（C语言实现）的源码。通过查看这个文件夹，可以了解到Go语言的方方面面。
~~~
查看：
```bash
chao@chao-PC:/usr/local/go/src$ ls
all.bash          clean.bat  errors    iostest.bash   os         sort
all.bat           clean.rc   expvar    log            path       strconv
all.rc            cmd        flag      make.bash      plugin     strings
androidtest.bash  cmp.bash   fmt       make.bat       race.bash  sync
archive           compress   go        Make.dist      race.bat   syscall
bootstrap.bash    container  hash      make.rc        reflect    testing
bufio             context    html      math           regexp     text
buildall.bash     crypto     image     mime           run.bash   time
builtin           database   index     naclmake.bash  run.bat    unicode
bytes             debug      internal  nacltest.bash  run.rc     unsafe
clean.bash        encoding   io        net            runtime    vendor
```
### 7、test文件夹
~~~test
    存放测试Go语言自身代码的文件。通过阅读这些测试文件，可大致了解Go语言的一些特性和使用方法。
~~~
## 二、gopath工作区目录结构
~~~
    在环境变量中除了$GOPATH这样的显式变量外，Go语言还有两个隐含的环境变量——GOOS和GOARCH。
    GOOS代表程序构建环境的目标操作系统，可笼统地理解为Go语言安装到的那个操作系统的标识，其值可以是darwin、freebsd、linux或windows。
    GOARCH则代表程序构建环境的目标计算架构，可笼统地理解为Go语言安装到的那台计算机的计算架构的标识，其值可以是386、amd64或arm。
~~~
> 工作区有3个子目录：src目录、pkg目录和bin目录。

### 1、src目录
~~~src
    用于以代码包的形式组织并保存Go源码文件。这里的代码包，与src下的子目录一一对应。例如，若一个源码文件被声明为属于代码包logging，那么它就应当被保存在src目录下名为logging的子目录中。当然，我们也可以把Go源码文件直接放于src目录下，但这样的Go源码文件就只能被声明为属于main代码包了。除非用于临时测试或演示，一般还是建议把Go源码文件放入特定的代码包中。
~~~
> Go语言的源码文件分为3类：Go库源码文件、Go命令源码文件和Go测试源码文件。

### 2、pkg目录
~~~pkg
    用于存放经由go install命令构建安装后的代码包（包含Go库源码文件）的“.a”归档文件。该目录与GOROOT目录下的pkg功能类似。区别在于，工作区中的pkg目录专门用来存放用户（也就是程序开发者）代码的归档文件。构建和安装用户源码的过程一般会以代码包为单位进行，比如logging包被编译安装后，将生成一个名为logging.a的归档文件，并存放在当前工作区的pkg目录下的平台相关目录中。
~~~

### 3、bin目录
~~~
 与pkg目录类似，在通过go install命令完成安装后，保存由Go命令源码文件生成的可执行文件。在Linux操作系统下，这个可执行文件一般是一个与源码文件同名的文件。而在Windows操作系统下，这个可执行文件的名称是源码文件名称加.exe后缀。
~~~
> 注意: 这里有必要明确一下Go语言的命令源码文件和库源码文件的区别。所谓命令源码文件，就是声明为属于main代码包，并且包含无参数声明和结果声明的main函数的源码文件。  
这类源码文件可以独立运行（使用go run命令），也可被go build或go install命令转换为可执行文件。而库源码文件则是指存在于某个代码包中的普通源码文件。

## 三、go编译时，目录查找顺序
> go工程包含依赖包管理，GOROOT，GOPATH三类目录来查找编译需要的库。他们的顺序如下：

1. 从工程项目的root目录查找vendor目录中的依赖库。
2. 从用户环境变量$GOPATH/src中查找依赖库。
3. 从用户环境变量$GOROOT/src中查找依赖库。
4. 未找到，抛出异常，编译终止。

## 总结：
~~~
    通过对golang的目录结构的了解和编译时查找依赖库的顺序，对这门语言有一个初步的认识，接下来我们将通过go的内部命令深入了解一下它。
~~~

