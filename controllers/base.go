package controllers

import (
	"fmt"
	"net/http"
	m "simple_web/models"
	u "simple_web/utils"
	"strconv"
)

//CreateItem - store one more item in storage :)
var CreateItem = func(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Create")
	var name string = r.URL.Query().Get("item_name")
	id, err := strconv.ParseUint(r.URL.Query().Get("item_id"), 10, 32)
	if err != nil {
		u.Message(false, "ID parse error")
	}
	var itemType string = r.URL.Query().Get("item_type")
	//	item := &m.Item{Name: r.Context().Value("item_name").(string), ID: r.Context().Value("item_id").(uint), Type: r.Context().Value("item_type").(string)}
	item := &m.Item{Name: name, ID: uint(id), Type: itemType}

	stored, err := item.Store()
	resp := u.Message(true, "success")
	resp["item"] = stored
	u.Respond(w, resp)
}

var DeleteItem = func(w http.ResponseWriter, r *http.Request) {
}

var GetItem = func(w http.ResponseWriter, r *http.Request) {
}

var ListItems = func(w http.ResponseWriter, r *http.Request) {
}
