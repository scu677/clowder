/*package main

import (
//"fmt"
"net/http"
"text/template"
)
type webCounter struct{
	count chan int
	template *template.Template
}
func NewCounter() *webCounter{
	counter := new(webCounter)
	counter.count = make(chan int, 1)
	go func(){
		for i:=1 ;; i++ { counter.count <- i }	
	}()
	return counter
	
}
func(w *webCounter) ServeHTTP(r http.ResponseWriter, rq *http.Request){
	if rq.URL.Path != "/"{
		r.WriteHeader(http.StatusNotFound)
		return
	}
	w.template.Execute(r, struct{Counter int}{<-w.count})
}
*/
