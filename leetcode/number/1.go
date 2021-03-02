package number
/*
1、两数之和

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
你可以按任意顺序返回答案


1、两数之和Ⅱ - 输入有序数组
2、两数之和Ⅲ - 数据结构设计
3、两数之和Ⅳ - 输入 BST
4、三数之和
5、四数之和
*/

func twoSum(nums []int, target int) []int {
	 m:=make(map[int]int)
	 for i,v:=range nums{
		if index,ok:=m[target-v];ok{
			return []int{index,i}
		}else {
			m[v]=i
		}
	 }
	return []int{}
}