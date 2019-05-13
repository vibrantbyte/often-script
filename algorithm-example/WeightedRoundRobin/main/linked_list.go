/*
@Time : 2019-04-12 18:47 
@Author : xiaoyueya
@File : linked_list
@Software: GoLand
*/
package main

//Node
type Node struct {
	Name string
	//链表
	Last *Node
	Next *Node
	Previous *Node
}
