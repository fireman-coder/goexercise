package divide

func myPow(x float64, n int) float64 {

	if n<0{
		return 1/devide(x,n)
	}
	return devide(x,n)
}

func devide (x float64,n int)float64{
	if n==0{
		return 1
	}
	y:=devide(x,n/2)
	if n%2==0{
		return y*y
	}
	return y*y*x
}
