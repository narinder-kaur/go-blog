package mysql

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

func Find(table string, wherecondition map[string]interface{}, selectfields string, options map[string]interface{}, holder interface{})([]interface{}){
	var conn *sql.DB
	conn, err := connect();
	checkErr(err)
	
	query := "select "+ selectfields + " from " + table + "where "
	attached_and := false
	and_clause := ""
	for key, value := range wherecondition {
		if(attached_and == true){
			and_clause = " and "
		}
		query += and_clause + key + "='" + value.(string) + "'" 
		attached_and = true
	}

	if limit, ok := options["limit"]; ok {
	 query += "limit "+ limit.(string)
	}


	result, err := conn.Query(query)
	checkErr(err)
	var collection []interface{}
	for result.Next() {
            result.Scan(&holder)
            collection = append(collection,holder)
            
	}
		return collection
}

func FindOne(table string, wherecondition map[string]interface{}, selectfields string, options map[string]interface{}, holder interface{})(interface{}){
	var conn *sql.DB
	conn, err := connect();
	checkErr(err)
	
	query := "select "+ selectfields + " from " + table + "where "
	attached_and := false
	and_clause := ""
	for key, value := range wherecondition {
		if(attached_and == true){
			and_clause = " and "
		}
		query += and_clause + key + "='" + value.(string) + "'" 
		attached_and = true
	}
	query += "limit 1"

	result, err := conn.Query(query)
	checkErr(err)
	for result.Next() {
            result.Scan(&holder)
            return holder
	}
		return holder
}

func Save(table string, data map[string]interface{}) int64 {
	// query = "insert into "+table+"(username, password, email) values(?,?,?)", u.Name,u.Password,u.Email
// func Save(query string, data ...string) int64{
	// var conn *sql.DB
	// conn, err := connect();
	// checkErr(err)
	
	// stmt, err := conn.Prepare(query)
	// checkErr(err)

	// result, err := stmt.Exec(data[0], data[1], data[2])
	// checkErr(err)

	// id, err := result.LastInsertId()
	// checkErr(err)

	// return id
	return 0
}

func checkErr(err error){
	if (err != nil){
		panic(err)
	}
}
