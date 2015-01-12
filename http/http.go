package http

import (
	"fmt"
	"github.com/denniscollective/dragonfly.go/dragonfly"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func StartServer() error {
	r := mux.NewRouter()
	r.HandleFunc("/favicon.ico", hollar)
	r.HandleFunc("/{b64JobString}", serveFile)

	err := http.ListenAndServe(":2345", handle(r))
	return err
}

type benchmarkHandler struct {
	handler http.Handler
}

func (h benchmarkHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t := time.Now()
	h.handler.ServeHTTP(w, req)
	fmt.Println(time.Now().Sub(t))
}

func handle(h http.Handler) http.Handler {

	return benchmarkHandler{handlers.CombinedLoggingHandler(os.Stdout, h)}

}

func hollar(response http.ResponseWriter, r *http.Request) {
	fmt.Println("hollar")
}

func serveFile(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(request)
	jobStr := vars["b64JobString"]
	file, err := dragonfly.ImageFor(jobStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, _ := ioutil.ReadAll(file)
	response.Write(data)
}
