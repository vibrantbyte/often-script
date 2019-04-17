/*
@Time : 2019-04-17 11:34 
@Author : xiaoyueya
@File : user_tree.go
@Software: GoLand
*/
package main

import (
	"fmt"
)

//树大小
var TreeSize int32 = 1024

var Tree *GTree

//地图用户树
type GTree struct {
	//节点存储集合
	GNodes []*GNode
	//根的位置下标和结点数
	R,N int32
}


// --- begin ---

//GNode
type GNode struct {
	//数据域
	Data interface{}
	//双亲在数组中的下标
	Parent int32
	//当前节点在数组中的下标
	Index int32

	//FirstChild
	FirstChild *GChild
}

//GChild
type GChild struct {
	Child int32
	Next *GChild
}


//GetParent
func (node *GNode) GetParent() *GNode {
	if node.Parent >= Tree.N || node.Parent == -1{
		return nil
	}
	return Tree.GNodes[node.Parent]
}

//GetIndex
func (node *GNode) GetIndex() int32{
	return node.Index
}

//Add
func (node *GNode) Add(item *GNode) *GNode{
	Tree.R ++
	Tree.GNodes[Tree.R] = item
	item.Index = Tree.R
	item.Parent = node.Index
	//计算总量
	Tree.N++

	//操作子节点
	tempNode := node.FirstChild
	if tempNode == nil {
		child := new(GChild)
		child.Child = item.Index
		child.Next = nil
		node.FirstChild = child
		return item
	}

	var tempChild *GChild
	for {
		if tempNode == nil {
			child := new(GChild)
			child.Child = item.Index
			child.Next = nil
			tempChild.Next = child
			break
		}else{
			tempChild = tempNode
			tempNode = tempNode.Next
		}
	}
	return item
}

//GetChild
func (node *GNode) GetChild() []*GNode {
	nodes := make([]*GNode,0)
	child := node.FirstChild
	for {
		if child == nil {
			break
		}
		nodes = append(nodes,Tree.GNodes[child.Child])
		child = child.Next
	}
	return nodes
}

// --- end ---






func main()  {

	root := new(GNode)
	root.Data = "北京"
	Tree = CreateTree(root)


	area1 := new(GNode)
	area1.Data = "区域1"

	area2 := new(GNode)
	area2.Data = "区域2"

	root.Add(area1)
	root.Add(area2)


	partnerId1 := new(GNode)
	partnerId1.Data = "合伙人1"

	partnerId2 := new(GNode)
	partnerId2.Data = "合伙人2"

	area1.Add(partnerId1)
	area1.Add(partnerId2)


	var i int32
	for i = 0;i<Tree.N ;i++ {
		node := Tree.GNodes[i]
		var pData = "nil"
		if node.GetParent() != nil {
			pData = fmt.Sprint(node.GetParent().Data)
		}

		var childString = ""
		childs := node.GetChild()
		if childs!= nil {
			for _,item := range childs {
				childString += fmt.Sprint(item.Data)
				childString += ","
			}
		}

		fmt.Printf("%s-%d,父节点：%s,子节点：%s \n",fmt.Sprint(node.Data),node.Index,pData,childString)




	}

	fmt.Println(Tree.N)

}

//创建几点
func CreateTree(root *GNode) *GTree{
	Tree = new(GTree)
	Tree.GNodes = make([]*GNode,TreeSize)
	Tree.GNodes[0] = root
	Tree.R = 0
	Tree.N = 1

	//节点默认值
	root.Index = 0
	root.Parent = -1
	return Tree
}

func Testqq()  {



}
