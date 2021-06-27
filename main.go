package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fmt.Println("hello golang")

	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/hello",helloHandler)
	log.Fatal(http.ListenAndServe(":9999",nil))
}
// indexHandler
// test command :curl http://localhost:9999/
func indexHandler(resp http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(resp,"url.path = %q\n",req.URL.Path)
}
// helloHandler
// test command :curl http://localhost:9999/hello
func helloHandler(resp http.ResponseWriter, req *http.Request)  {
	for k, v := range req.Header {
		fmt.Fprintf(resp,"header[%q]=%q\n",k,v)
	}
}