package DBInsert

import (
	"testing"
    "fmt"
    //"log"
     "database/sql"
    _ "github.com/go-sql-driver/mysql"
    _"github.com/mattn/go-sqlite3"
)
    
    
    
type testservers struct {
	servername string 
	availability bool
	user string
	
}
	
var tests = []testservers{
	{"Test", false, "Tester"},
	{"Test2", true, "Tester2"},
}
	
func TestIncertAttribute(t *testing.T) {
  
  	db, err := sql.Open("sqlite3", "/usr/home/jdawson/goWorkspace/src/github.com/musec/clowder/design/Database-mock/TestDB.db")
      
    if err != nil {
        fmt.Println(err)
        fmt.Println("error one tripped")
        return 
    }
    defer db.Close()
    err = db.Ping()
    if err != nil {
        fmt.Println(err)
        fmt.Println("error Two tripped")
        return 
    }
    fmt.Println("connection established")
  
  
   for _, servers := range tests{
	   got := IncertAttribute(servers.servername, db)
	   if servers.availability != got{
		   t.Errorf("GetAvailblity() == %t  should have been  %t\n" , got, servers.availability)
	   }
   }
}
