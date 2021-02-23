package recursion

import (
	"strconv"
	"strings"
)

type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root==nil{
		return "#"
	}
	return strconv.Itoa(root.Val)+","+this.serialize(root.Left)+","+this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	list:=split(data)
	return build(&list)
}

func build(list *[]string)*TreeNode{
	if len(*list)==0{
		return nil
	}
	val:=(*list)[0]
	*list=(*list)[1:]
	if val=="#"{
		return nil
	}
	v,_:=strconv.Atoi(val)
	root:=&TreeNode{Val: v}
	root.Left=build(list)
	root.Right=build(list)
	return root
}

func split(data string)[]string{
	return strings.Split(data,",")
}


