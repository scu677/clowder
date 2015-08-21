package main

import (
    "database/sql"
    "log"
    "fmt"
    //_ "github.com/go-sql-driver/mysql"
    _"github.com/mattn/go-sqlite3"
  //  "strconv"
)

func main() {
    var table string = "tablename"

    db, err := sql.Open("sqlite3", TestDB+"/"+jdawson+"/"+abc123)//("sqlite3", "jdawson:@/TestDB")//"mymysql", database+"/"+user+"/"+password)  // does not work
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	checkConnection()
	
    // read data from database
    read(db, table)
}

func read(db *sql.DB, table string) {
    rows, err := db.Query("SELECT Availability FROM serverInformation WHERE Hostname='Test'", db)
	if err != nil{
		log.Fatal(err)
	}
	
	fmt.Println(rows)
}


func checkConnection(){
		//database.Ping()
	
	}
