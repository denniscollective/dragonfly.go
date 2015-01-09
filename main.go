package main

import (
	"fmt"
	"github.com/denniscollective/dragonfly.go/dragonfly"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{b64JobString}", serveFile)

	err := http.ListenAndServe(":2345", r)
	if err != nil {
		panic(err)
	}

}

func serveFile(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	jobStr := vars["b64JobString"]
	file, err := dragonfly.ImageFor(jobStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Name())

	data, _ := ioutil.ReadAll(file)
	response.Write(data)
}
