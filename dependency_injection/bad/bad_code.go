package bad

//bad case

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

// SomeFunc  业务函数，调用selectValFromDBByID函数
func SomeFunc(user User) bool {
	u, err := selectValFromDBByID(user.id)
	if err != nil {
		fmt.Println("selectValFromDBByID EXEC err: ", err)
		return false
	}

	//以下为一些业务逻辑
	fmt.Println(u)
	time.Sleep(time.Second)
	return true
}

// selectValFromDBByID 通过id查询数据并且返回
func selectValFromDBByID(id int) (User, error) {
	//以下是对DB查询操作，只是Demo，没有加入三方驱动，也没有实际的数据，不可以运行

	// dsn DB地址
	dsn := "user:password@tcp(127.0.0.1:3306)/DBName"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close() // 先判断是否连接异常，再进行关闭连接

	sqlStr := "select id, name, age from user where id=?"
	var u User
	err = db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("QueryRow EXEC err:%v\n", err)
		return User{}, err
	}
	return u, nil
}
