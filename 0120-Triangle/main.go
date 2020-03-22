package main

import (
	"fmt"
)

func minimumTotal1(triangle [][]int) int {

	lens := len(triangle)

	for i := 1; i < len(triangle); i++ {
		for j := 0; j <= i; j++ {
			switch {
			case j == 0:
				triangle[i][j] += triangle[i-1][j]
			case i == j:
				triangle[i][j] += triangle[i-1][j-1]
			case triangle[i-1][j-1] < triangle[i-1][j]:
				triangle[i][j] += triangle[i-1][j-1]
			default:
				triangle[i][j] += triangle[i-1][j]
			}
		}
	}
	res := triangle[lens-1][0]

	for i := 0; i < lens; i++ {
		if res > triangle[lens-1][i] {
			res = triangle[lens-1][i]
		}
	}

	return res
}

func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	path := make([][]int, n)
	for i := 0; i < n; i++ {
		path[i] = make([]int, len(triangle[i]))
	}

	path[0][0] = triangle[0][0]

	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			path[i][j] = 99999999
			// 这里处理边界处理的比较恶心...
			if j == 0 {
				path[i][j] = path[i-1][j] + triangle[i][j]
			} else if i == j {
				path[i][j] = path[i-1][j-1] + triangle[i][j]
			} else {
				path[i][j] = min(path[i-1][j], path[i-1][j-1]) + triangle[i][j]
			}
		}
	}

	res := 9999999
	for i := 0; i < n; i++ {
		res = min(res, path[n-1][i])
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal3(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}

	return triangle[0][0]
}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	res1 := minimumTotal1(triangle)
	res2 := minimumTotal2(triangle)
	res3 := minimumTotal3(triangle)

	fmt.Println(res1, res2, res3)
}
