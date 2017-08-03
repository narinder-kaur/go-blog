package main

import (
	"net/http"
	`first_web_app/models`
	"fmt"
)

type signup string
type saveuser string

func (c signup) ServeHTTP(res http.ResponseWriter, req *http.Request){
	
	view("signup",nil, res)
}

func (c saveuser) ServeHTTP(res http.ResponseWriter, req *http.Request){
	req.ParseForm()
	u := models.User{Name:req.PostFormValue("name"),Email:req.PostFormValue("email"),Password:req.PostFormValue("password")}
	if ok, err := u.Validate(); (ok == true && err == nil) {
		fmt.Println("saving");
		fmt.Println(u.Save());
	}
}

func main(){
	var s signup
	var save saveuser 
	m := http.NewServeMux()
	m.Handle("/",s)
	m.Handle("/save-user",save)

	http.ListenAndServe(":8002",m)
}