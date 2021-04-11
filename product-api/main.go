package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return
		}
		// log.Printf("Data %s\n", d)
		fmt.Fprintf(rw, "Hello %s", d)
	})
	http.HandleFunc("/goodbey", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbey")
	})
	http.ListenAndServe(":9090", nil)
}
