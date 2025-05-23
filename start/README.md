* 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
   
    https://github.com/776727681/goTask/blob/master/main01.go 
* 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。
  
   https://github.com/776727681/goTask/blob/master/main02.go
  
* 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效

  https://github.com/776727681/goTask/blob/master/main03.go

* 查找字符串数组中的最长公共前缀
  
  https://github.com/776727681/goTask/blob/master/main04.go

* 给定一个排序数组，你需要在原地删除重复出现的元素 【给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。】

  https://github.com/776727681/goTask/blob/master/main05.go

* 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一

  https://github.com/776727681/goTask/blob/master/main06.go

* 删除有序数组中的重复项 【给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。】

  https://github.com/776727681/goTask/blob/master/main07.go

* 合并区间 【以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中】

  https://github.com/776727681/goTask/blob/master/main08.go

* 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数

  https://github.com/776727681/goTask/blob/master/main09.go