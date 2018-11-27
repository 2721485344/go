package main
import ( "database/sql"
          "fmt"
            _ "github.com/go-sql-driver/mysql" )
func main() {
	db, _ := sql.Open("mysql", "root:2018@/ms")
	defer db.Close()
	rows, _ := db.Query("select * from tb_user;")
	defer rows.Close()
	cloumns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(cloumns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i] }
	for rows.Next() {
		rows.Scan(scanArgs...)
	    var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
				} else {
					value = string(col)
					}
			fmt.Println(cloumns[i], " --- ", value)
		}
		}
	//----------------------------------增删改----------------------
	db3, _ := sql.Open("mysql", "root:2018@/ms")
	defer db3.Close()
	stmt, _ := db3.Prepare("INSERT INTO ms VALUES ('MS-1GOGOGO')")
	stmt.Exec()

	}
