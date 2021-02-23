package recursion
var resarr [][]int
//46. 全排列，给定一个 没有重复 数字的序列，返回其所有可能的全排列。
func permute(nums []int) [][]int {
	resarr=make([][]int,0)
	generate2(nums,make([]int,0))
   return resarr
}

func generate2(nums []int,demo []int){

	if len(nums)==0{
		resarr=append(resarr,demo)
		return
	}

	for i,_:=range nums{
		//demo1:=make([]int,len(demo))
		//copy(demo1,demo)
		demo1:=append(demo,nums[i])
		//append(nums[:index],nums[index+1:]...) 这里千万不能用，会改变nums指针内容，而len，cap不变，遍历时出问题
		generate2(deleteIndex(nums,i),demo1)
	}
}

func deleteIndex(nums []int,index int )[]int{
	nums1:=make([]int,len(nums))
	copy(nums1,nums)  //关键
	return append(nums1[:index],nums1[index+1:]...)
}