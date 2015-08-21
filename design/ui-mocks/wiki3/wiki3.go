// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	//"github.com/musec/clowder/design/DBandUI/DBInsert"
	"github.com/musec/clowder/design/DBandUI/DBGet"
	//"github.com/musec/clowder/design/DBandUI/DBConnect"
	
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    _"github.com/mattn/go-sqlite3"
)

var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)

type Page struct {
	Title string
	Body  []byte
	Hostname string
	User string
	Availblity string
} 

type testservers struct {
	servername string 
	availability bool
	BookedUntill, user, MacAddress, IP, processor, cores, RAM, DiskSpace,NetCard,Comments string
	
}
	
var tests = []testservers{
	{"Test", false, "a", "Tester", "b", "123.123.123.123", "i27",  "c", "d", "e", "f", "g"},
	{"Test2", true, "1", "Tester2", "2", "3", "4", "5", "6", "7", "8", "9"},
}

	var db *sql.DB
	var err error

/***********************************************************************
 * main
 * 
***********************************************************************/

func main() {
	/******************************************************************/
	db, err = makeConnection("/usr/home/jdawson/goWorkspace/src/github.com/musec/clowder/design/Database-mock/TestDB.db")

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
	user := DBGet.GetUser("Test", db)
	fmt.Println(user)
/******************************************************************/

	
	
	
	
	flag.Parse()
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	if *addr {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
		if err != nil {
			log.Fatal(err)
		}
		s := &http.Server{}
		s.Serve(l)
		return
	}

	http.ListenAndServe(":8080", nil)
}


func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}


/***********************************************************************
 * loadPage: loads page based on the title it is passed
 * Return: page, error
***********************************************************************/
func loadPage(title string) (*Page, error) {
	PageInfo := new(Page)
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	PageInfo.Body = body
	//PageInfo.Hostname = hostname
	
	PageInfo.Title = title
	//user := "error"							//important
	if db != nil{
		PageInfo.User = DBGet.GetUser("Test", db)
		PageInfo.Availblity = DBGet.GetAvailblity("Test" ,db)
	}else{
		PageInfo.User = "error db = nil"
	}
	if err != nil {
		return nil, err
	}
	//return &Page{Title: title, Body: body, User: user}, nil//, TableEntry: tableEntry
	return PageInfo, nil
}


/***********************************************************************
 * viewHandler
 * 
***********************************************************************/
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
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




