package main

import "flag"

var (
	h bool
	v bool
	n string
)

// init 赋值初始化项目
func init() {
	flag.BoolVar(&h, "h", false, "this help.")
	flag.BoolVar(&v, "v", false, "show version.")
	flag.StringVar(&n, "name", "", "show project name.")

	flag.Usage = func() {
		flag.PrintDefaults()
	}
}

//主函数接收flag信息
func main() {
	flag.Parse()

	if h {
		flag.Usage()
	}
}
