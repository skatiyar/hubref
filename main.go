package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
