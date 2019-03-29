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