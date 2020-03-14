package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	if len(s) <= 0 {
		return ""
	}

	// 把字符串读到map里面
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	var countArr []int
	for _, v := range m {
		countArr = append(countArr, v)
	}

	// 这个语法真丑
	sort.Sort(sort.Reverse(sort.IntSlice(countArr)))

	res := ""
	for _, v := range countArr {
		for k := range m {
			if v == m[k] {
				fmt.Println(v, m[k], k)
				t := ""
				for i := 0; i < v; i++ {
					t += string(k)
				}
				res += t
				delete(m, k)

			}
		}
	}

	return res
}

func main() {
	frequencySort("tree")
}
