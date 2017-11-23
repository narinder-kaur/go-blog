package main

import (
	"fmt"
	
	"net/http"
	`first_web_app/models` 
	"github.com/alexedwards/scs"
	
)

type signup string
type saveuser string
type homepage string
type login string
type logout string
type newblog string


var sessionManager = scs.NewCookieManager("u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4")



func (c signup) ServeHTTP(res http.ResponseWriter, req *http.Request){
	view("signup",nil, res)
}


func (h homepage) ServeHTTP(res http.ResponseWriter, req *http.Request){
	session := sessionManager.Load(req)
	user, _ := session.GetString("user")
    
	if (len(user) > 0){
		view("dashboard",user, res)	
	}else{
		view("homepage",nil, res)
	}
}
func (c newblog) ServeHTTP(res http.ResponseWriter, req *http.Request){
	if(req.Method == http.MethodGet){
		view("newblog",nil, res)
	}

	if(req.Method == http.MethodPost){
		fmt.Println(req.PostFormValue("title"));
		fmt.Println(req.PostFormValue("description"));
		fmt.Println(req.PostFormValue("image"));
		fmt.Println(req.PostFormValue("tags"));
	}
}

func (c logout) ServeHTTP(res http.ResponseWriter, req *http.Request){
 	session := sessionManager.Load(req)
	user, _ := session.GetString("user")
	if (len(user) > 0){
		session.Destroy(res)
	}
	http.Redirect(res,req,"/",http.StatusFound)
}

func (c login) ServeHTTP(res http.ResponseWriter, req *http.Request){
var user models.User
	req.ParseForm()
	data := map[string]string{"email":req.PostFormValue("email"),"password":req.PostFormValue("password")}
	u,_ := user.Find(data)
	if (u.Uid >0){
				session := sessionManager.Load(req)
			// defer sess.SessionRelease(res)
			session.PutString(res,"user", u.Name)
			http.Redirect(res,req,"/",http.StatusFound)
	}
		http.Redirect(res,req,"/",http.StatusFound)

}
func (c saveuser) ServeHTTP(res http.ResponseWriter, req *http.Request){
	
	var user models.User
	req.ParseForm()
	u := models.User{Name:req.PostFormValue("name"),Email:req.PostFormValue("email"),Password:req.PostFormValue("password")}
	if ok, err := u.Validate(); (ok == true && err == nil) {
		if user, err = u.Save(); err!= nil{
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
			session := sessionManager.Load(req)
			// defer sess.SessionRelease(res)
			session.PutString(res,"user", user.Name)
			http.Redirect(res,req,"/",http.StatusFound)
		}
	}


func main(){
	var s signup
	var l login
	var logout logout
	var save saveuser 
	var home homepage
	var newblog newblog
	m := http.NewServeMux()
	m.Handle("/signup",s)
	m.Handle("/save-user",save)
	m.Handle("/",home)
	m.Handle("/login",l)
	m.Handle("/logout",logout)
	m.Handle("/create-new",newblog)
	m.Handle("/save-blog",newblog)
	http.ListenAndServe(":8002",m)
}