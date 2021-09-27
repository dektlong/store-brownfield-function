package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"encoding/json"
 	"fmt"
 	"io/ioutil"
)

var API_CALL="sample-api"

func handler(w http.ResponseWriter, r *http.Request) {
	
	log.Println(r.RemoteAddr, r.Method, r.URL.String())
	
	fmt.Fprintf(w, "<H1>Welcome to app-name function</H1>")
		    
	executeAPI (API_CALL)
	
	log.Println("Function revision:", os.Getenv("REV"))
	
}

func executeAPI (string apiCall) {
	response, err := http.Get(apiCall)

    	if err != nil {
        	log.Println(err.Error())
        	os.Exit(1)
    	}

    	responseData, err := ioutil.ReadAll(response.Body)
	
    	if err != nil {
        	log.Fatal(err)
		os.Exit(1)
    	}
    	
	fmt.Fprintf(w,string(responseData))
}
	
func main() {
	
	http.HandleFunc("/", handler)

	var addr = flag.String("addr", ":8080", "addr to bind to")
	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
