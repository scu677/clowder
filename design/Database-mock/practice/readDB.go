/**********************************************************************
 *  
 * 
 * Reference pages
 * Package sql:   http://golang.org/pkg/database/sql/
 * Package sql Trotroial: http://go-database-sql.org/index.html
 * 
 * 
 * 
 **********************************************************************/






package DBaccess


import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	
)



func main() {
	db, err := sql.Open("sqlite3",
		"/usr/home/jdawson/goWorkspace/src/github.com/musec/clowder/design/Database-mock/TestDB.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	err = db.Ping()
		if err != nil {
	// do something here
	}
	
}	
	
/*	
	var (
	id int
	name string
)
rows, err := db.Query("select id, name from users where id = ?", 1)
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	err := rows.Scan(&id, &name)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id, name)
}
err = rows.Err()
if err != nil {
	log.Fatal(err)
}



stmt, err := db.Prepare("select id, name from users where id = ?")
if err != nil {
	log.Fatal(err)
}
defer stmt.Close()
rows, err := stmt.Query(1)
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	// ...
}
if err = rows.Err(); err != nil {
	log.Fatal(err)
}





var name string
err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
if err != nil {
	log.Fatal(err)
}
fmt.Println(name)




stmt, err := db.Prepare("select name from users where id = ?")
if err != nil {
	log.Fatal(err)
}
var name string
err = stmt.QueryRow(1).Scan(&name)
if err != nil {
	log.Fatal(err)
}
fmt.Println(name)

	
	
	
}


*/
