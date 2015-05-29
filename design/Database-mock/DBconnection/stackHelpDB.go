/***********************************************************************
*
* 
***********************************************************************/



package main

import (
    "database/sql"
    "fmt"
    "log"
    "strings"
    _ "github.com/go-sql-driver/mysql"
    _"github.com/mattn/go-sqlite3"
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
	isAvailable := GetAvailblity(servername, db)
	fmt.Println(isAvailable) 
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
  
/***********************************************************************
* this gets the Availblity of the server from the database
* Returns: bool true if it is in use
* 
***********************************************************************/

func GetAvailblity(serverHostName string, Database *sql.DB) bool { 
	isAvailable := GetAttribute(serverHostName, Database, "Availability")
	if isAvailable == "true" || isAvailable == "True"{
		return true
	}else{
		return false
	}
}


/***********************************************************************
* this gets the current user of the server from the database
* Returns: string of the user. 
* 
***********************************************************************/
func GetUser(serverHostName string, Database *sql.DB) string {   
	
	return GetAttribute(serverHostName, Database, "User")
}

/***********************************************************************
* this quearies the databace based on what attrbuite it was passed in
* Returns: string of data in the 
* 
***********************************************************************/
func GetAttribute(serverHostName string, Database *sql.DB, attribute string) string {   
	
    rows, err := Database.Query(strings.Replace("SELECT ? FROM serverInformation WHERE Hostname=?", "?", attribute, 1), serverHostName)
    if err != nil {
            log.Fatal(err)
    }
    var user string
    defer rows.Close()
    for rows.Next() {
            if err := rows.Scan(&user); err != nil {
                    log.Fatal(err)
            }
            fmt.Printf("curent User is %s for %s\n", user, serverHostName)
    }
    if err := rows.Err(); err != nil {
            log.Fatal(err)
    }  
    return user
}
