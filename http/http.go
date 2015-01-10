package http

import (
	"fmt"
	"github.com/denniscollective/dragonfly.go/dragonfly"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func StartServer() error {
	r := mux.NewRouter()
	r.HandleFunc("/favicon.ico", hollar)
	r.HandleFunc("/{b64JobString}", serveFile)

	err := http.ListenAndServe(":2345", r)

	return err
}

func hollar(response http.ResponseWriter, r *http.Request) {
	fmt.Println("hollar")
}

func serveFile(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("%s %s %s", request.RemoteAddr, request.Method, request.URL)
	response.Header().Set("Access-Control-Allow-Origin", "*")

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
