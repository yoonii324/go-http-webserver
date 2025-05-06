package main

import (
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

func main() {
	//http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":3000", nil)
}
