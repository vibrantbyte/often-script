package main

import "fmt"

// 定义矩阵
var sourceMatrix [3][5]int

// init 给sourceMatrix赋值
func init() {
	sourceMatrix = [3][5]int{
		{1, 2, 0, 3, 4},
		{2, 3, 4, 5, 1},
		{1, 1, 5, 3, 0},
	}
}

// main 求一个矩阵中最大的二维矩阵（各元素之和最大）
func main() {
	//定义二维数组
	var targetMatrix [2][2]int

	//遍历数组
	//定义循环列
	column := 4
	//定义循环行
	row := 2
	//定义缓存最大值
	var maxNum int

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			temp := sourceMatrix[i][j] + sourceMatrix[i+1][j] + sourceMatrix[i][j+1] + sourceMatrix[i+1][j+1]
			if temp > maxNum {
				maxNum = temp
				//将坐标记录一下
				targetMatrix[0][0] = sourceMatrix[i][j]
				targetMatrix[1][0] = sourceMatrix[i+1][j]
				targetMatrix[0][1] = sourceMatrix[i][j+1]
				targetMatrix[1][1] = sourceMatrix[i+1][j+1]
			}
		}
	}

	fmt.Print(targetMatrix)

}
