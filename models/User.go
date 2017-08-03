package models

import(
"reflect"
"fmt"
_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string
	Email string
	Password string
}

func (u User) Save () (User, error){
	return u,nil
}

func (u User) Validate() (bool,error){
	v := reflect.ValueOf(u)
	fmt.Println(v)
	return true,nil
}

