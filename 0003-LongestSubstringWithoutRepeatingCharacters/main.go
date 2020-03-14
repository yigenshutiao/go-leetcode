package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

	var (
		maxLens = 0
		curLens = 0
	)

	var d = map[byte]bool{}

	for i := 0; i < len(s); i++ {
		// 如果字符在里面
		if d[s[i]] == true {
			l := i - curLens
			for l < i {
				delete(d, s[l])
				curLens -= 1

				if s[l] == s[i] {
					break
				}
				l += 1
			}
		}
		// 如果字符不在里面
		d[s[i]] = true
		curLens += 1
		if curLens > maxLens {
			maxLens = curLens
		}
	}
	return maxLens
}

func main() {
	s := "abcbad"
	res := lengthOfLongestSubstring(s)
	fmt.Println("length is :", res)
}
