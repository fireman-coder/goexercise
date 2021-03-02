package vote

/*
给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。

所有出现超过 ⌊ n/3 ⌋ 次的元素-->最多有2个元素
摩尔投票法：
cond1，count1
cond2，count2
抵消阶段：
1、如果遍历的数不是cond1，cond2，那么都-1；
2、如果是其中一个（cond2），那么cond2++
3、cond1，cond2之间不产生冲突
计数阶段：
count1>0,count2>0,则可能是
*/

func majorityElement_229(nums []int) []int {
	var res []int
	cond1,cond2,count1,count2:=nums[0],nums[0],0,0

	for _,v:=range nums{
		if cond1==v{
			count1++
			continue
		}

		if cond2==v{
			count2++
			continue
		}

		if count1==0{
			cond1=v
			count1++
			continue
		}

		if count2==0{
			cond2=v
			count2++
			continue
		}

		count1--
		count2--
	}

	if count1!=0 && counttime(nums,cond1)>len(nums)/3{
		res=append(res,cond1)
	}

	if count2!=0&& counttime(nums,cond2)>len(nums)/3{
		res=append(res,cond2)
	}
	return res
}

func counttime(nums []int,targrt int)int {
	count:=0
	for _,v:=range nums{
		if v==targrt{
			count++
		}
	}
	return count
}