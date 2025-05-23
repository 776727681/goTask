package main

import "fmt"

func main()  {
	strs := []string{"flower","flow","flight"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)

}


func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	oneStr := strs[0]
	for i := 0; i < len(oneStr); i++ {
		for _, s := range strs[1:] {
			if s[i] != oneStr[i] {
				return oneStr[:i]
			}
		}
	}
	return oneStr
}