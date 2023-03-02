package good

//本着最小依赖的原则，不要使用大而全的接口，这样在 mock 的时候就能减轻不少工作量，这和接口隔离原则不谋而合

type IBasicOperate interface {
	SetInitVal(int)
	Result() int
}

type IAddSubOperate interface {
	IBasicOperate
	Add(int)
	Subtract(int)
}

type IMulDivOperate interface {
	IBasicOperate
	Multiply(int)
	Divide(int)
}

type IComplex interface {
	IAddSubOperate
	IMulDivOperate
}

// AddSubOperate 基本计算：只包含加减
type AddSubOperate struct {
	IAddSubOperate
	InitNum int
	AddNum  int
	SubNum  int
}

func (basic AddSubOperate) Computer() int {
	basic.SetInitVal(basic.InitNum)
	basic.Add(basic.AddNum)
	basic.Subtract(basic.SubNum)
	return basic.Result()
}

// Complex 复杂计算：包含加减乘除
type Complex struct {
	IComplex
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
