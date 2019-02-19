package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL" //_匿名导入的方式
)

type company struct {
	id      int
	name    string
	name_en string
}

/**数据库连接*/
func main() {

	//sql.Open并不会立即建立一个数据库的网络连接, 也不会对数据库链接参数的合法性做检验, 它仅仅是初始化一个sql.DB对象. 当真正进行第一次数据库查询操作时, 此时才会真正建立网络连接;
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/db1")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	query(db)

	id, err := insert(db)

	query(db)

	update(db, id)

	query(db)

	del(db, id)

	query(db)
}

func insert(db *sql.DB) (int64, error) {

	fmt.Println("insert ---")
	result, err := db.Exec("insert into company(name, name_en) values (?, ?)", "go", "go语言")

	fmt.Println(result, err)

	if err == nil {
		return result.LastInsertId()
	}

	return -1, err
}

func update(db *sql.DB, id int64) {

	fmt.Println("update ---")
	result, err := db.Exec("update company set name_en = ? where id = ?", "go language", id)

	fmt.Println(result, err)
}

func del(db *sql.DB, id int64) {

	fmt.Println("del ---")
	result, err := db.Exec("delete from company where id = ? ", id)

	fmt.Println(result, err)
}

func query(db *sql.DB) {
	fmt.Println("select ---")
	//查询
	rows, err := db.Query("select * from company ")
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var comp company
		rows.Scan(&comp.id, &comp.name, &comp.name_en)
		fmt.Println(comp)
	}
}
