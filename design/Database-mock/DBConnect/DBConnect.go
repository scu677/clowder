package main

import (
    "database/sql"
    "fmt"
    //"log"
    //"strings"
    _ "github.com/go-sql-driver/mysql"
    _"github.com/mattn/go-sqlite3"
	"github.com/musec/clowder/design/Database-mock/DBInsert"
	"github.com/musec/clowder/design/Database-mock/DBGet"   
)
func main() {


	db, err := makeConnection("/usr/home/jdawson/goWorkspace/src/github.com/musec/clowder/design/Database-mock/TestDB.db")

	defer db.Close()
	err = db.Ping()
	if err != nil {
        fmt.Println(err)
        fmt.Println("error Two tripped")
        ///ToDo Add reconnect code
        db.Close()
        return 
    }  
        
	servername:= "Test"
	isAvailable := DBGet.GetAvailblity(servername, db)
	fmt.Println(isAvailable) 
	
	DBInsert.UpdateProcessor(servername, db, "i27")
	DBInsert.UpdateIP(servername, db, "123.123.123.123")
	
	
	return

}
/***********************************************************************
* this makes the Database connection
* Returns: reference to the Database, any errors Recived
* 
***********************************************************************/
func makeConnection(dbDirectory string) (*sql.DB, error) {
	 db, err := sql.Open("sqlite3", dbDirectory)
      
    if err != nil {
        fmt.Println(err)
        fmt.Println("error one tripped")
        return db, err
    }
    //defer db.Close()
    err = db.Ping()
    if err != nil {
        fmt.Println(err)
        fmt.Println("error Two tripped")
        return db, err
    }
    fmt.Println("connection established")
	
	
	return db, err
   
  }
