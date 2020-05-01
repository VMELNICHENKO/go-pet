package main

import (
	"fmt"
	"log"
	"net/http"
	c "simple_web/controllers"
)

func main() {
	fmt.Println("Start listening")
	//	router.HandleRoutes()

	http.HandleFunc("/create", c.CreateItem)
	http.HandleFunc("/delete", c.DeleteItem)
	http.HandleFunc("/get", c.GetItem)
	http.HandleFunc("/list", c.ListItems)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Print("Hello") })
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("LintenAndServe: ", err)
	}
}
