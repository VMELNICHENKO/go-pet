package router

import (
	"fmt"
	"net/http"
	c "simple_web/controllers"
)

// HandleRoutes - simple redirection to controller
func HandleRoutes() {
	mux := http.NewServeMux()
	ch := http.HandlerFunc(c.CreateItem)
	dh := http.HandlerFunc(c.DeleteItem)
	gh := http.HandlerFunc(c.GetItem)
	ah := http.HandlerFunc(c.ListItems)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Print("Hello") })
	mux.Handle("/create", ch)
	mux.Handle("/delete", dh)
	mux.Handle("/get", gh)
	mux.Handle("/list", ah)
}

//RouterHandler - point of requests handling
// func RouterHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	fmt.Println(r.Form)
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Fuck off!")
// }
