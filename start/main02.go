package main

import (
	"fmt"
	"strconv"
)

func main()  {
	x := 121
	res := isPalindrome(x)
	fmt.Println(res)
}

func isPalindrome(x int) bool {
	if x <= 0 {
		return false
	}
	str1 := strconv.Itoa(x)
	var str2 string
	for i := 0; i < len(str1); i++ {
		str2 = string(str1[i]) + str2
	}
	if str1 == str2 {
		return true
	}
	return false
}
