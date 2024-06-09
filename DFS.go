package main

import (
	"fmt"
)

func dfsUtils(mat [][]int, vis [][]int, i int, j int, rows int, cols int, target int) bool{
	if i<0 || j<0 || i>=rows || j>=cols || mat[i][j]==0 || vis[i][j]==1 {
		return false
	}

	if target == mat[i][j] {
		return true
	}

	return dfsUtils(mat, vis, i+1, j, rows, cols, target) ||
	       dfsUtils(mat, vis, i, j+1, rows, cols, target) ||
		   dfsUtils(mat, vis, i-1, j, rows, cols, target) ||
		   dfsUtils(mat, vis, i, j-1, rows, cols, target);
}

func dfs(mat [][]int, rows int, cols int, target int) {
	vis := make([][]int, rows)
	for i:=0; i<rows; i++ {
		vis[i]=make([]int, cols)
		for j:=0; j<cols; j++ {
			vis[i][j]=0
		}
	}

	for i:=0; i<rows; i++ {
		for j:=0; j<cols;j++ {
			if mat[i][j]==1 {
				if dfsUtils(mat, vis, i, j, rows, cols, target) {
					fmt.Print("\nPath exists")
					return;
				}
			}
		}
	}
	fmt.Print("\nElement does not exist")
}

func main() {
	var rows, cols, target int
	fmt.Print("\nEnter the no. of rows:")
	fmt.Scan(&rows)
	fmt.Print("\nEnter the no. of cols:")
	fmt.Scan(&cols)

	mat := make([][]int, rows)

	for i:=0; i<rows; i++ {
		mat[i] = make([]int, cols)
	}

	fmt.Print("\nEnter the matrix:")
	for i:=0; i<rows; i++ {
		for j:=0; j<cols; j++ {
			fmt.Scan(&mat[i][j])
		}
	}

	fmt.Print("\nEnter the target:")
	fmt.Scan(&target)

	dfs(mat, rows, cols, target)
}

