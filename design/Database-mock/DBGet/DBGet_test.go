package DBGet

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
	BookedUntill, user, MacAddress, IP, processor, cores, RAM, DiskSpace,NetCard,Comments string
	
}
	
var tests = []testservers{
	{"Test", false, "a", "Tester", "b", "123.123.123.123", "i27",  "c", "d", "e", "f", "g"},
	{"Test2", true, "1", "Tester2", "2", "3", "4", "5", "6", "7", "8", "9"},
}	



type fullRow struct {
	servername string 
	availability bool
	BookedUntill, user, MacAddress, IP, processor, cores, RAM, DiskSpace,NetCard,Comments string
	
}
	
var allRows = []fullRow{
	{"Test", false, "a", "Tester", "b", "123.123.123.123", "i27",  "c", "d", "e", "f", "g"},
	{"Test2", true, "1", "Tester2", "2", "3", "4", "5", "6", "7", "8", "9"},
}
/*
var allRows = *sql.Rows{
	{"Test", false, "a", "Tester", "b", "123.123.123.123", "i27",  "c", "d", "e", "f", "g"},
	{"Test2", true, "1", "Tester2", "2", "3", "4", "5", "6", "7", "8", "9"},
}	*/
	var db *sql.DB
	var err error
	
	
/***********************************************************************
* this function is not a test, it is used to make the database connection
* Returns:
* 
***********************************************************************/

func dbconnectTest(){
	db, err = sql.Open("sqlite3", "/usr/home/jdawson/goWorkspace/src/github.com/musec/clowder/design/Database-mock/TestDB.db") 
    if err != nil {
        fmt.Println(err)
        fmt.Println("error one tripped")
        return 
    }
   //defer db.Close()
    err = db.Ping()
    if err != nil {
        fmt.Println(err)
        fmt.Println("error Two tripped")
        return 
    }
    fmt.Println("connection established")
    return 
}


func TestGetUser(t *testing.T) {
    dbconnectTest()
    fmt.Println("GetUser")
    
	for _, servers := range tests{
	   got := GetUser(servers.servername, db)
	   if servers.user != got{
		   t.Errorf("GetUser Failed GetAvailblity() == %s  should have been  %s\n" , got, servers)
	   }
   }	
}

func TestGetAvailblity(t *testing.T) {

	fmt.Println("TestGetAvailblity")  
  
   for _, servers := range tests{
	   got := GetAvailblity(servers.servername, db)
	   if servers.availability != got{
		   t.Errorf("GetAvailblity() == %t  should have been  %t\n" , got, servers.availability)
	   }
   } 
}
/*
func TestGetRow(t *testing.T){
	
    fmt.Println("TestGetRow")
  
	for _, servers := range tests{
		//fmt.Println("TestGetRow for ran")
		got := GetRow(servers.servername, db)
		gots := RowToTests(got)
		//fmt.Println(gots)
		if servers != gots{
			fmt.Println(gots)
			t.Errorf("GetAvailblity() == %s  should have been  %t\n" , gots, servers.availability)
		}
	}   	
}*/


func TestGetRow(t *testing.T){
	
    fmt.Println("TestGetRow")
  
	for _, servers := range tests{
		//fmt.Println("TestGetRow for ran")
		got := GetRow(servers.servername, db)
		gots := RowToTests(got)
		//fmt.Println(gots)
		if servers != gots{
			fmt.Println(gots)
			t.Errorf("GetAvailblity() == %s  should have been  %t\n" , gots, servers.availability)
		}
	}   	
}

/*
func TestGetAllRow(t *testing.T){
	
    fmt.Println("TestGetAllRow")
  
   for _, servers := range tests{
		got := GetAllRows(servers.servername, db)
		for got.Next() {
			fmt.Println(got)
		}
		//gots := RowToTests(got)
		//fmt.Println(gots, "for ran")
		//if servers != allRows{
		//	t.Errorf("GetAvailblity() == %s  should have been  %s\n" , gots, servers)
		//}
	}   
}
*/
func TestGetEntryCount(t *testing.T){
	
	fmt.Println("TestGetEntryCount")
  
	for _, servers := range tests{
		got := GetEntryCount(db)
		fmt.Println(got)
		if servers.user != got{
			t.Errorf("GetAvailblity() == %t  should have been  %t\n" , got, servers.availability)
		}
	}
	db.Close()
}



//try reflection 
func RowToTests(RecivedRow Row)testservers{
	var RowTest testservers
	RowTest.servername = RecivedRow.servername  
	RowTest.availability = RecivedRow.availability
	RowTest.BookedUntill = RecivedRow.BookedUntill
	RowTest.user = RecivedRow.user
	RowTest.MacAddress = RecivedRow.MacAddress
	RowTest.IP = RecivedRow.IP
	RowTest.processor = RecivedRow.processor
	RowTest.cores = RecivedRow.cores
	RowTest.RAM = RecivedRow.RAM
	RowTest.DiskSpace = RecivedRow.DiskSpace
	RowTest.NetCard = RecivedRow.NetCard
	RowTest.Comments = RecivedRow.Comments

	return RowTest
}

/*
func RowToTests(RecivedRow *sql.Rows)testservers{
	var RowTest testservers
	RowTest.servername = RecivedRow.servername  
	RowTest.availability = RecivedRow.availability
	RowTest.BookedUntill = RecivedRow.BookedUntill
	RowTest.user = RecivedRow.user
	RowTest.MacAddress = RecivedRow.MacAddress
	RowTest.IP = RecivedRow.IP
	RowTest.processor = RecivedRow.processor
	RowTest.cores = RecivedRow.cores
	RowTest.RAM = RecivedRow.RAM
	RowTest.DiskSpace = RecivedRow.DiskSpace
	RowTest.NetCard = RecivedRow.NetCard
	RowTest.Comments = RecivedRow.Comments

	return RowTest
}*/

