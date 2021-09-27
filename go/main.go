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
		    
	fmt.Fprintf(w, "API: ", API_CALL)
	fmt.Fprintf(w, "<BR>Response:<BR>")
	
	response, err := http.Get(API_CALL)

    	if err != nil {
        	log.Println(err.Error())
        	fmt.Fprintf(w,"Unable to exectute this API")
    	}

    	responseData, err := ioutil.ReadAll(response.Body)
	
    	 if err != nil {
		log.Println(err.Error())
		fmt.Fprintf(w,"Unable to exectute this API")
    	} else {
		fmt.Fprintf(w,string(responseData))
	}
	
	log.Println("Function revision:", os.Getenv("REV"))
	
}

func executeAPI (string apiCall) {
	
}
	
func main() {
	
	http.HandleFunc("/", handler)

	var addr = flag.String("addr", ":8080", "addr to bind to")
	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
