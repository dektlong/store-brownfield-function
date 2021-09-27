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
	
	log.Println("Calling API: ", apiCall)
	
	client := &http.Client{}
 	
	req, err := http.NewRequest("GET", apiCall, nil)
 	
	if err != nil {
  		log.Println(err.Error())
 	}
	
 	req.Header.Add("Accept", "application/json")
 	req.Header.Add("Content-Type", "application/json")
 	resp, err := client.Do(req)
 	
	if err != nil {
  		log.Println(err.Error())
 	}
	
	defer resp.Body.Close()
 	bodyBytes, err := ioutil.ReadAll(resp.Body)
 	
	if err != nil {
  		log.println(err.Error())
 	}

	var responseObject Response
 	json.Unmarshal(bodyBytes, &responseObject)
 	log.Println("API Response as struct %+v\n", responseObject)
}
	
func main() {
	
	http.HandleFunc("/", handler)

	var addr = flag.String("addr", ":8080", "addr to bind to")
	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
