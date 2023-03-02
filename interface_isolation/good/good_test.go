package good

import (
	"fmt"
	"testing"
)

// 打桩：实现 IOperate 接口
type basicOP struct {
	num int
}

func (o *basicOP) SetInitVal(i int) {
	o.num = i
}

func (o *basicOP) Add(i int) {
	o.num += i
}

func (o *basicOP) Subtract(i int) {
	o.num -= i
}

func (o *basicOP) Result() int {
	return o.num
}

// complexOp 实现IComplex
type complexOp struct {
	basicOP
}

func (c *complexOp) Divide(i int) {
	//除0处理：打印提示改除法操作无效，然后把除数设为1
	if i == 0 {
		fmt.Printf("Divide invalid")
		i = 1
	}
	c.num /= i
}

func (c *complexOp) Multiply(i int) {
	c.num *= i
}

func TestBasicCase(t *testing.T) {
	op := new(basicOP)
	basic := AddSubOperate{op, 1, 2, 3}
	// 1+2-3 = 0
	ret := basic.Computer()
	if 0 != ret {
		t.Errorf("expect 0,  actually%d\n", ret)
	}
}

func TestComplexCase(t *testing.T) {
	op := new(complexOp)
	complex := Complex{op, 1, 12, 3, 4, 5}
	// (1+12-3)*4/5 = 0
	ret := complex.Computer()
	if ret != 8 {
		t.Errorf("expect 8,  actually%d\n", ret)
	}

	//测试除0的情况
	complex0 := Complex{op, 1, 12, 3, 4, 0}
	// (1+12-3)*4/1 = 40
	ret = complex0.Computer()
	if ret != 40 {
		t.Errorf("expect 40,  actually%d\n", ret)
	}
}
