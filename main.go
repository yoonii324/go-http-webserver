package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// func barHandler(w http.ResponseWriter, r *http.Request) {
// 	values := r.URL.Query()    // bring query parameters
// 	name := values.Get("name") // check if there are any specific key values
// 	if name == "" {
// 		name = "World"
// 	}

// 	id, _ := strconv.Atoi(values.Get("id")) // get the id value and cast it to an integer type
// 	fmt.Fprintf(w, "Hello %s! id: %d", name, id)
// }

// func main() {
// 	mux := http.NewServeMux() // create ServeMux instance
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello World") // register handler to the instance
// 	})
// 	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello Bar")
// 	})

// 	// http.HandleFunc("/bar", barHandler) // register bar handler
// 	http.ListenAndServe(":3000", mux) // start mux instance
// }

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHandler() http.Handler { // function that makes handler instance
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"aaa", 16, 87}
	data, _ := json.Marshal(student)                   // translate Student object into []byte format
	w.Header().Add("content-type", "application/json") // denote that it is JSON format
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data)) // send the result
}

// func main() {
// 	//http.Handle("/", http.FileServer(http.Dir("static")))
// 	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

// 	http.ListenAndServe(":3000", MakeWebHandler())
// }
