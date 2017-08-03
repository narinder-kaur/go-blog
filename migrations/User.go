package main


import (
  _ "github.com/go-sql-driver/mysql"
 "first_web_app/core"
 "fmt"
)

func main(){
 
 queries := []string{
	 "CREATE TABLE userinfo (uid INT(10) NOT NULL AUTO_INCREMENT, username VARCHAR(64) NOT NULL, password VARCHAR(64) NOT NULL,email VARCHAR(250) NOT NULL, PRIMARY KEY (uid))",
	}

	for _,query := range queries {
	result := core.Execute(query)
	fmt.Println(result);
	}
}