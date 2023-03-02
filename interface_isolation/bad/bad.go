package bad

type IOperate interface {
	SetInitVal(int)
	Add(int)
	Subtract(int)
	Multiply(int)
	Divide(int)
	Result() int
}

// Basic 基本计算：只包含加减
type Basic struct {
	IOperate
	InitNum int
	AddNum  int
	SubNum  int
}

func (b Basic) Computer() int {
	b.SetInitVal(b.InitNum)
	b.Add(b.AddNum)
	b.Subtract(b.SubNum)
	return b.Result()
}

// Complex 复杂计算：包含加减乘除
type Complex struct {
	IOperate
	InitNum int
	AddNum  int
	SubNum  int
	MulNum  int
	DivNum  int
}

func (c Complex) Computer() int {
	c.SetInitVal(c.InitNum)
	c.Add(c.AddNum)
	c.Subtract(c.SubNum)
	c.Multiply(c.MulNum)
	c.Divide(c.DivNum)
	return c.Result()
}
