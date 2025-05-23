package main

import (
	"fmt"
)

func main()  {
	nums := []int{2,7,11,15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}
func twoSum(nums []int, target int) []int {
	if len(nums) == 0{
		return nil
	}
	hmap := make(map[int]int)
	for i:=0; i<=len(nums); i++{
		target1 := target - nums[i]
		if t_k ,ok := hmap[target1];ok{//判断target1在map中key是否存在
			return []int{t_k,i}
		}
		// 将当前数值和索引存入哈希表
		hmap[nums[i]] = i
	}
	// 题目保证有解，此返回仅为语法需要
	return nil
}
