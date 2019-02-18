package main

import (
	"errors"
	"fmt"
)

/**结构体**/
type IInfo interface {
	info()
}

/**	结构体 */
type User struct {
	name string
	password string
	age int
}

type Member struct {
	name string
	mobile string
}

func (user User) info()  {
	fmt.Printf("user name is %s \n", user.name)
}

func (member Member) info() error {
	if member.name == "" { return errors.New(" null value ") }

	fmt.Printf("member name is %s \n", member.name)

	return nil
}

func main() {

	var user User
	user.name = "zhang3"
	user.password = "zhang3 pwd"
	user.age = 33

	fmt.Println(user.name)
	fmt.Println(user)

	li4 := User{"li4", "li4 pwd", 44}
	fmt.Println(li4)

	/* 结构体指针 */
	fmt.Println(&user)
	up := &user
	fmt.Println(up)
	fmt.Println(up.name)

	user.info()

	var member Member
	r := member.info()
	fmt.Println(r)
	member = Member{"wang2", "13111118888"}
	member.info()
	
}
