package main

import (
	"fmt"
	"github.com/gorilla/mux"
)

type user  struct{
	name string
	numberOfVisits int
}
func GetName(name string)(string,int){
	type users map[string]int
	for username, numberOfVisits := range users{
		if username == name{
			return username,numberOfVisits
		}
	}
	users["name"]=1
	return name,1
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{name}", GetName)
	http.Handle("/", r)
}
