package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	body, readErr := ioutil.ReadAll( r.Body)
	for _, hed := range r.Header{
		fmt.Println(hed)
	}

	str := string(body)
	if readErr != nil {
		log.Fatal(readErr)
	}
//js,_ := json.Marshal(str)
	fmt.Printf(str)
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "host" : "localhost" ,"person" : {"firstname": "John", "lastname" :"Doe"}}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(str))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
