# golang 生成 shared object 供脚本语言使用

# LINUX so 文件基本概念和命名规则
![版本示意图](./images/go-shared-object.png)

**libxmns.so.1.2.3 1 major 2 minor 3 release**

* major 增加，原有函数接口已经不能使用，minor和release 复归于0
* minor 增加， 新增加了一些函数接口，但原有函数接口还能使用， release 复归于0
* release 增加，修改一些bug, 函数接口不变

# 模板

```go
package main

import "C"
import "fmt"

//export Sum
func Sum(a int, b int) int {
	return a + b
}

//export GetName
func GetName(firstName string) string{
	return fmt.Sprint(firstName,"-so")
}

func main(){

}
```
> 注意，即使是要编译成动态库，也要有main函数，上面的import "C"一定要有 而且一定要有注释

# 编译
```bash
go build -buildmode=c-shared -o libhello.so .\libhello.go
```