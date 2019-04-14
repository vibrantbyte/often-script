/*
@Time : 2019-04-14 09:57 
@Author : xiaoyueya
@File : strategy
@Software: Goand
*/
package main

import "fmt"

type A interface {
	GetName() string
}

type A1 struct {
	Name string
}

func (a *A1) GetName() string  {
	a.Name = "ceshi"
	return a.Name
}

type A2 struct {
	A1
}

func main()  {

	var a A = new(A2)
	TestType(a)


}

func TestType(a A){
	if item,ok := a.(*A1);ok{
		fmt.Println(item.GetName())
	}

}





