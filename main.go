package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()    // bring query parameters
	name := values.Get("name") // check if there are any specific key values
	if name == "" {
		name = "World"
	}

	id, _ := strconv.Atoi(values.Get("id")) // get the id value and cast it to an integer type
	fmt.Fprintf(w, "Hello %s! id: %d", name, id)
}

func main() {
	http.HandleFunc("/bar", barHandler) // register bar handler
	http.ListenAndServe(":3000", nil)   // start a web server
}
