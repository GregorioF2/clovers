package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gorilla/mux"
	config "github.com/gregorioF2/clovers/configs"
	middlewares "github.com/gregorioF2/clovers/middlewares"
	riddles "github.com/gregorioF2/clovers/routes/riddles"
)

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/riddles/jug", riddles.JugRiddleHandler)
	router.Use(middlewares.SetCorsHeaders)
	return router
}

func main() {
	router := newRouter()
	fmt.Printf("Server listening on port %s\n", config.SERVER_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.SERVER_PORT), router))
}
