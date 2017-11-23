package models

import(
// "reflect"
"fmt"
db "first_web_app/core/mongodb"
)

type User struct {
	Uid int64
	Name string
	Email string
	Password string
}

func (u User) Save () (User, error){
	data = map[string]interface{}{
		"username": u.Name,
		"email":u.Email,
		"password":u.Password
	}
	result := db.Save("userinfo", data)
		u.Uid = result
		return u, nil
}

func (u User) Validate() (bool,error){
	// v := reflect.ValueOf(u)
	return true,nil
}

func (u User) Find(condition map[string]string) (User, error){
	query := "select * from userinfo where "
	
	query += "Limit 1"
	result, _ := db.FindOne("userinfo",condition,"*",map[string][int]{"limit":1}, User)
	
	return result,nil
	
}

