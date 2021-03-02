package recall
/*
78. 子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/
var res [][]int
func subsets(nums []int) [][]int {
	res=make([][]int,0)
	recall(nums,make([]int,0),len(nums))
	return res
}

func recall( nums,demo []int,count int ){
	if count==0{
		//res=append(res,demo) 不能直接append，因为demo是引用类型，放入res后，可能会被重新修改
		tmp := make([]int, len(demo))
		copy(tmp, demo)
		res=append(res,tmp)
		return
	}
	if len(nums)!=0{
		recall(nums[1:],demo,count-1)
	}

	for i:=0;i<len(nums);i++{
		//demo_bak:=copy1(demo,nums[i])
		demo=append(demo,nums[i])
		recall(nums[i+1:],demo,count-1)
		demo=demo[:len(demo)-1]
	}

}

func copy1(demo []int,val int)[]int{
	demo_bak:=make([]int,len(demo))
	copy(demo_bak,demo)
	demo_bak=append(demo_bak,val)
	return demo_bak
}
