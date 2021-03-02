package recall

var phoneMap map[string][]string=map[string][]string{
	"2":{"a","b","c"},
	"3":{"d","e","f"},
	"4":{"g","h","i"},
	"5":{"j","k","l"},
	"6":{"m","n","o"},
	"7":{"p","q","r","s"},
	"8":{"t","u","v"},
	"9":{"w","x","y","z"},
}
var result []string
func letterCombinations(digits string) []string {
	result=make([]string,0)
	if len(digits)==0{
		return result
	}
	backtrack(digits,"",0)
	return result
}

func backtrack(digits ,combination string,count int){
	if len(digits)-1==count{
		result=append(result,combination)
		return
	}
	letter:=phoneMap[string(digits[count])]
	count++
	for _,v:=range letter{
		backtrack(digits,combination+v,count)
	}
}



//func main(){
//	str:="1234"
//	for _,v:=range strings.Split(str,""){
//		fmt.Println(v) //1,2,3,4
//	}
//}