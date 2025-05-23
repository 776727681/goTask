package main

import "fmt"

func main()  {
	digits := []int{4,3,2,1}
	res := plusOne(digits)
	fmt.Println(res)
}
func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
		if carry == 0 {
			break
		}
	}
	if carry != 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}