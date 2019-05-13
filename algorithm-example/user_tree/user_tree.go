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

//getParent
func (node *GNode) getParent(tree *GTree) *GNode {
	if node.Parent >= tree.N || node.Parent == -1{
		return nil
	}
	return tree.GNodes[node.Parent]
}

//GetIndex
func (node *GNode) GetIndex() int32{
	return node.Index
}

//Add
func (node *GNode) insert(tree *GTree,item *GNode) *GNode{
	tree.R ++
	tree.GNodes[tree.R] = item
	item.Index = tree.R
	item.Parent = node.Index
	//计算总量
	tree.N++

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

//getChild
func (node *GNode) getChild(tree *GTree) []*GNode {
	nodes := make([]*GNode,0)
	child := node.FirstChild
	for {
		if child == nil {
			break
		}
		nodes = append(nodes,tree.GNodes[child.Child])
		child = child.Next
	}
	return nodes
}

//TreeSize
var TreeSize int32 = 4096

//地图用户树
type GTree struct {
	//节点存储集合
	GNodes []*GNode
	//根的位置下标和结点数
	R,N int32
	//树大小
	TreeSize int32
}

//DefaultCreatTree
func DefaultCreatTree() *GTree{
	tree := new(GTree)
	tree.GNodes = nil
	tree.R = -1
	tree.N = 0
	tree.TreeSize = TreeSize
	return tree
}

//CreateTree
func CreateTree(root *GNode) *GTree{
	tree := new(GTree)
	tree.GNodes = make([]*GNode,TreeSize)
	tree.GNodes[0] = root
	tree.R = 0
	tree.N = 1
	tree.TreeSize = TreeSize

	//节点默认值
	root.Index = 0
	root.Parent = -1
	return tree
}

//SetTreeSize
func (tree *GTree) SetTreeSize(treeSize int32){
	if tree.N > 0 || tree.GNodes != nil {
		panic("设置树大小必须在插入节点之前！")
	}
	tree.TreeSize = treeSize
}

//ClearTree
func (tree *GTree) ClearTree(){
	tree.N = 0
	tree.R = -1
	tree.GNodes = nil
}

//TreeEmpty
func (tree *GTree) TreeEmpty() bool {
	if tree.GNodes == nil && tree.N == 0 {
		return true
	}else {
		return false
	}
}

//Root
func (tree *GTree) Root() *GNode {
	if tree.TreeEmpty() {
		return nil
	}else{
		return tree.GNodes[0]
	}
}

//Value
func (tree *GTree) Value(cur *GNode) interface{}{
	if cur == nil {
		return nil
	}else{
		return cur.Data
	}
}

//Assign
func (tree *GTree) Assign(cur *GNode,value interface{}){
	if cur != nil{
		cur.Data = value
	}
}

//Parent
func (tree *GTree) Parent(cur *GNode) *GNode{
	if cur == nil {
		return nil
	}
	return cur.getParent(tree)
}

//InsertChild
func (tree *GTree) InsertChild(node *GNode,cur *GNode){
	//树默认大小
	if tree.N >= tree.TreeSize{
		panic("树的节点已经到最大值不能进行插入！")
	}
	if tree.GNodes == nil && tree.N == 0 {
		tree.GNodes = make([]*GNode,tree.TreeSize)
		tree.GNodes[0] = cur
		tree.R = 0
		tree.N = 1

		//节点默认值
		cur.Index = 0
		cur.Parent = -1
		return
	}
	if node != nil {
		node.insert(tree,cur)
	}
}

//DeleteChild
func (tree *GTree) DeleteChild(cur *GNode){

}

//ChildList
func (tree *GTree) ChildList(node *GNode) []*GNode {
	if node == nil {
		return nil
	}
	return node.getChild(tree)
}

//Length
func (tree *GTree) Length() int32{
	return tree.N
}





func main()  {
	tree := DefaultCreatTree()
	tree.SetTreeSize(60)

	root := new(GNode)
	root.Data = "北京"
	tree.InsertChild(nil,root)


	area1 := new(GNode)
	area1.Data = "区域1"

	area2 := new(GNode)
	area2.Data = "区域2"

	tree.InsertChild(root,area1)
	tree.InsertChild(root,area2)


	partnerId1 := new(GNode)
	partnerId1.Data = "合伙人1"

	partnerId2 := new(GNode)
	partnerId2.Data = "合伙人2"

	tree.InsertChild(area1,partnerId1)
	tree.InsertChild(area1,partnerId2)


	var i int32
	for i = 0;i<tree.N ;i++ {
		node := tree.GNodes[i]
		var pData = "nil"
		if tree.Parent(node) != nil {
			pData = fmt.Sprint(tree.Parent(node).Data)
		}

		var childString = ""
		childs := tree.ChildList(node)
		if childs!= nil {
			for _,item := range childs {
				childString += fmt.Sprint(tree.Value(item))
				childString += ","
			}
		}

		fmt.Printf("%s-%d,父节点：%s,子节点：%s \n",fmt.Sprint(node.Data),node.Index,pData,childString)




	}

	fmt.Println(tree.N)

}


func Testqq()  {



}
