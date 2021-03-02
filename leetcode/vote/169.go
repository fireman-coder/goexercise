package vote

/*
摩尔投票法：
candidate（候选人），count（统计票数）
抵消阶段：
1、遍历时，候选人一致，count+1，否则-1
2、遍历a时，发现count==0时，candidate=1
计数阶段：
最后maj=candidate且count>0

注意：
1、最后得到的抵消票数不为 0 的话，那说明他可能希望的
2、如果存在>n/2次的数，那就是他
*/

/*
169. 多数元素，多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
次数>n/2,肯定是最多的，否则就>n了
*/

func majorityElement(nums []int) int {
	candidate,count:=0,0
	for _,v:=range nums{
		if count==0{
			candidate=v
		}

		if v==candidate{
			count++
		}else {
			count--
		}
	}
	if count<1{
		return -1
	}
	//即使count>1，candidate也不一定大于n/2，比如[2,2,3,3,1];建议在遍历统计下candidate，判断是否>n/2
	//最后得到的抵消票数不为 0 的话，那说明他可能希望的
	return candidate
}




