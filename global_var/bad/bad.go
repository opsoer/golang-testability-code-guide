package bad

var GlobalNum int

func GlobalNumAdd(i int) int {
	GlobalNum += i
	return GlobalNum
}

func GlobalNumSub(i int) int {
	GlobalNum -= i
	return GlobalNum
}
