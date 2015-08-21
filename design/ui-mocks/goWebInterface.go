package main 

import (
"fmt"
//"time"
"net/http"
)

func main(){
	
	http.HandleFunc("/", handler)
	
	http.HandleFunc("/earth", handler2)
	
	http.ListenAndServe(":8080", nil)
	
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello Earth\n")
}







/*

func count (id int){
	
	for i := 0; i<10; i++{
		fmt.Println(id, ":", i)
		
		time.Sleep(time.Sleep(time.Millisecond *1000)
	}
	
	
}

func main(){
		for i := 0; i<10; i++{
			go count(i)
		}
	
}
*/
