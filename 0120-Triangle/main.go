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

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	res := minimumTotal1(triangle)
	fmt.Println(res)
}
