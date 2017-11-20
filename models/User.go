package models

import(
// "reflect"
"fmt"
"first_web_app/core"
)

type User struct {
	Uid int64
	Name string
	Email string
	Password string
}

func (u User) Save () (User, error){
	result := core.Save("insert into userinfo(username, password, email) values(?,?,?)", u.Name,u.Password,u.Email)
		u.Uid = result
		return u, nil
}

func (u User) Validate() (bool,error){
	// v := reflect.ValueOf(u)
	return true,nil
}

func (u User) Find(condition map[string]string) (User, error){
	query := "select * from userinfo where "
	attached_and := false
	and_clause := ""
	for key, value := range condition {
		fmt.Println(key+value)
		if(attached_and == true){
			and_clause = " and "
		}
		query += and_clause + key + "='" + value + "'" 
		attached_and = true
	}
	query += "Limit 1"
	result, _ := core.Execute(query)
	for result.Next() {
            var uid int64
            var username string
            var password string
            var email string
            result.Scan(&uid, &username, &password, &email)
            return User{uid,username,email,password}, nil
            
        }
	return User{},nil
	
}

