package core

import( 
  _ "github.com/go-sql-driver/mysql"	
"database/sql"
)

func connect() (*sql.DB, error) {
	connection, _ := sql.Open("mysql", "root:Narinder*105@/go_blog?charset=utf8")
	if err := connection.Ping(); err != nil {
		return connection, err
		}
	return connection, nil
}

func read(){}

func Execute(query string) (*sql.Rows, error){
	var conn *sql.DB
	conn, err := connect();
	checkErr(err)

	result, err := conn.Query(query)
	checkErr(err)

	return result, nil
}


func Save(query string, data ...string) int64{
	var conn *sql.DB
	conn, err := connect();
	checkErr(err)
	
	stmt, err := conn.Prepare(query)
	checkErr(err)

	result, err := stmt.Exec(data[0], data[1], data[2])
	checkErr(err)

	id, err := result.LastInsertId()
	checkErr(err)

	return id
}

func checkErr(err error){
	if (err != nil){
		panic(err)
	}
}
