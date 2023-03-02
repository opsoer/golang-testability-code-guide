package good

import (
	"errors"
	"testing"
)

// 本test文件 即使我没有加入三方数据库引擎，没有搭建数据库server，我依然 可以运行下面代码。Good code！！！

// 自己实现 SomeFunc 需要的 DBOperate接口参数，然后需要时注入（inject）SomeFunc
type MyDB struct{}

func (db *MyDB) selectValFromDBByID(id int) (User, error) {
	if id == 1 {
		return User{id: 1, age: 21, name: "xiao fang 1"}, nil
	}
	if id == 2 {
		return User{id: 2, age: 22, name: "xiao fang 2"}, nil
	}

	if id == 3 {
		return User{id: 3, age: 23, name: "xiao fang 3"}, nil
	}
	return User{}, errors.New("not found")
}

func TestSomeFunc(t *testing.T) {
	db := &MyDB{}
	ok := SomeFunc(db, User{id: 1, age: 21, name: "xiao fang 1"})
	if !ok {
		t.Error("SomeFunc func err")
	}

	ok = SomeFunc(db, User{id: 2, age: 22, name: "xiao fang 2"})
	if !ok {
		t.Error("SomeFunc func err")
	}

	ok = SomeFunc(db, User{id: 3, age: 23, name: "xiao fang 3"})
	if !ok {
		t.Error("SomeFunc func err")
	}

	//边界测试 ok == false
	ok = SomeFunc(db, User{})
	if ok {
		t.Error("SomeFunc func err")
	}
	// Output:
	//=== RUN   TestSomeFunc
	//{1 21 xiao fang 1}
	//{2 22 xiao fang 2}
	//{3 23 xiao fang 3}
	//selectValFromDBByID EXEC err:  not found
	//--- PASS: TestSomeFunc (3.03s)

}
