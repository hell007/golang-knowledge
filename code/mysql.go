/**
 * name: 使用MySQL数据库
 * author: jie
 * note:
 *
	关键的几个函数我解释一下：

	sql.Open()函数用来打开一个注册过的数据库驱动，go-sql-driver中注册了mysql这个数据库驱动，第二个参数是DSN(Data Source Name)，它是go-sql-driver定义的一些数据库链接和配置信息。它支持如下格式：

	user@unix(/path/to/socket)/dbname?charset=utf8
	user:password@tcp(localhost:5555)/dbname?charset=utf8
	user:password@/dbname
	user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
	db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。

	db.Query()函数用来直接执行Sql返回Rows结果。

	stmt.Exec()函数用来执行stmt准备好的SQL语句

	我们可以看到我们传入的参数都是=?对应的数据，这样做的方式可以一定程度上防止SQL注入。
 *
*/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:admin@/jie?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("曹小曹", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("曹小刘", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}
