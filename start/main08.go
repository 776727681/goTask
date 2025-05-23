package main

import (
	"fmt"
	"sort"
)

func main()  {
	intervals := [][]int{{1,3},{2,6},{8,10},{15,18}}
	res := merge(intervals)
	fmt.Println(res)
}

func merge(intervals [][]int)  [][]int{
	//判断是否为nil
	if len(intervals) == 0{
		return intervals
	}
	//按区间起点排序
	// 对二维切片 intervals 进行 排序 排序规则是：按照每个子切片的 第一个元素（区间的起始点） 从小到大排列。
	//匿名函数比较第二层数组第一个元素  返回 true 表示 intervals[i] 应该排在 intervals[j] 前面。;回 false 表示 intervals[j] 应该排在 intervals[i] 前面。
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println("----------------")
	fmt.Println(intervals)
	merged := [][]int{intervals[0]}
	fmt.Println("-----------11-----")
	fmt.Println(merged)
	for _, current := range intervals[1:] {
		last := merged[len(merged)-1]// [1,3]
		if current[0] <= last[1] {
			// 合并区间，更新结束点
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
	
}



