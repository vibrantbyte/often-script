package main

import "fmt"

//GetName
func GetName() string {
	return "plugin-name"
}

func main(){
	fmt.Println(GetName())
}
