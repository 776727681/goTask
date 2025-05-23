package main

import "fmt"

func main()  {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	res := removeDuplicates(nums)
	fmt.Println(res)

}


func removeDuplicates(nums []int) int {
	i := 0
	for j:=0; j<len(nums);j++{
		if nums[j] != nums[i]{
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
