package main


import (
	// "time"
  _ "github.com/go-sql-driver/mysql"
 "first_web_app/core"
 "fmt"
 "strconv"
)

func main(){
 
 	queries := map[string]string{
		"userinfo":"CREATE TABLE userinfo (uid INT(10) NOT NULL AUTO_INCREMENT, username VARCHAR(64) NOT NULL, password VARCHAR(64) NOT NULL,email VARCHAR(250) NOT NULL, PRIMARY KEY (uid))",
		"blogs":"CREATE TABLE blogs (uid INT(10) NOT NULL AUTO_INCREMENT, title VARCHAR(1000) NOT NULL, description TEXT NOT NULL,image VARCHAR(255), PRIMARY KEY (uid))",

	}
	core.Execute("CREATE TABLE IF NOT EXISTS migrations (migration_name VARCHAR(255) NOT NULL UNIQUE, created_at bigint NOT NULL)")
	
	var mig string
	// var migration_query string
	for name,_ := range queries {
		 result := core.QueryOne("select migration_name from migrations where migration_name='"+name+"'",mig)
		  fmt.Println(strconv.Itoa(result))
		// 	fmt.Println("select * from migrations where migration_name='"+name+"'")
		// if(mig.migration_name != name){
		// 	// result, _ := core.Execute(query)
		// 	created_at := time.Now()
		// 	migration_query := fmt.Sprintf("Insert into migrations values('%s',%s)",name,created_at.Format("20060102150405"))
		// 	// core.Execute(migration_query)
		// 	fmt.Println("not run till now" + name + migration_query)
		// }else{
		// 	fmt.Println("Already run "+ name)
		// }
	}
}