/*
@Time : 2019-04-12 18:47 
@Author : xiaoyueya
@File : linked_list
@Software: GoLand
*/
package main

//Node
type Node struct {
	next *Node
	prev *Node
}

//NewNode
func NewNode(next,prev *Node) *Node{
	node := new(Node)
	return node
}
