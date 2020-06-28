package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Ooops"))
			return
		}
		// log.Printf("Data %s\n", d) // curl -v -d 'wafu' localhost:9090  go run main.go のログに Data wafu
		fmt.Fprintf(rw, "Hello %s", d) // curl のレスポンスに Hello wafu
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodby World")
	})
	http.ListenAndServe(":9090", nil)
}