package core

import( 
  _ "github.com/go-sql-driver/mysql"	
"database/sql"
"fmt"
)

func connect() *sql.DB {
	connection, _ := sql.Open("mysql", "root:Narinder*105@/go_blog?charset=utf8")
	return connection
}

func read(){}

func Execute(query string) *sql.Rows{
	result, err := connect().Query(query)
	if(err != nil){
		fmt.Println(err);
	}

	return result
}
