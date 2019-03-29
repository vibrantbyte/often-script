package main

import "C"

//export Sum
func Sum(a int, b int) int {
	return a + b
}

func main(){

}