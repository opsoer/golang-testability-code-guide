package good

/*
  good case：通过依赖注入来提高代码的可测试性和扩展性，不管底层dbServer改为MySQL、MongoDB、OceanBase或者其他的dbServer,
  我都不需要更改自己的业务逻辑代码。
*/

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	id   int
	age  int
	name string
}

// SomeFunc  业务函数
func SomeFunc(dbs DBOperate, user User) bool {
	u, err := dbs.selectValFromDBByID(user.id)
	if err != nil {
		fmt.Println("selectValFromDBByID EXEC err: ", err)
		return false
	}

	//以下为一些业务逻辑
	fmt.Println(u)
	time.Sleep(time.Second)
	return true
}

type DBOperate interface {
	selectValFromDBByID(id int) (User, error)
}

// DB 实现 DBOperate 接口，可以根据情况加一些字段
type DB struct{}

func NewDb() DBOperate {
	return &DB{}
}

// selectValFromDBByID 通过id查询数据并且返回
func (myDB *DB) selectValFromDBByID(id int) (User, error) {
	//以下是对DB查询操作，只是Demo，没有加入三方驱动，也没有实际的数据，不可以运行

	dsn := "user:password@tcp(127.0.0.1:3306)/DBName"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := "select id, name, age from user where id=?"
	var u User
	err = db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("QueryRow EXEC err:%v\n", err)
		return User{}, err
	}
	return u, nil
}
