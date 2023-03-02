package bad

import "testing"

// 看似UT很简单，但是你必须得 模拟一套数据库的环境，真要命！而且还得走网络，网络不确定因素太大了！bade code！！！
func TestSomeFunc(t *testing.T) {

	//弄三个类似的用例 是因为 可能需要测试一些边界条件
	ok := SomeFunc(User{id: 1, age: 21, name: "xiao fang 1"})
	if !ok {
		t.Error("SomeFunc func err")
	}

	ok = SomeFunc(User{id: 2, age: 22, name: "xiao fang 2"})
	if !ok {
		t.Error("SomeFunc func err")
	}

	ok = SomeFunc(User{id: 3, age: 23, name: "xiao fang 3"})
	if !ok {
		t.Error("SomeFunc func err")
	}

	ok = SomeFunc(User{})
	if ok {
		t.Error("SomeFunc func err")
	}
}
