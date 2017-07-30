package main

import (
	"net/http"
	"log"
)

type signup string
type saveuser string

func (c signup) ServeHTTP(res http.ResponseWriter, req *http.Request){
	
	view("signup",nil, res)
}

func (c saveuser) ServeHTTP(res http.ResponseWriter, req *http.Request){
	req.ParseForm();
	data := req.Form;
	for key,value := range data {
		
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