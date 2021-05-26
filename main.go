package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/julienschmidt/httprouter"
	"github.com/skatiyar/hubref/controllers"
	"github.com/skatiyar/hubref/database"
	m "github.com/skatiyar/hubref/middlewares"
)

var (
	port  = os.Getenv("PORT")
	user  = os.Getenv("USERNAME")
	pass  = os.Getenv("PASSWORD")
	dbURL = os.Getenv("DATABASE_URL")
)

func main() {
	if port == "" {
		log.Fatal("PORT must be set")
	}
	if dbURL == "" {
		log.Fatal("DATABASE_URL must be set")
	}
	if err := database.Connect(dbURL); err != nil {
		log.Fatal(err)
	}
	if err := database.CreateSchema(); err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	router.PanicHandler = dieHard

	// Routes
	router.GET("/", m.BasicAuth(controllers.Homepage, user, pass))
	router.GET("/api/paths", m.BasicAuth(m.AcceptJSON(controllers.GetPaths), user, pass))
	router.GET("/api/paths/*path", m.BasicAuth(m.AcceptJSON(controllers.GetPath), user, pass))
	router.POST("/api/paths", m.BasicAuth(m.AcceptJSON(controllers.CreatePath), user, pass))
	router.PUT("/api/paths/*path", m.BasicAuth(m.AcceptJSON(controllers.EditPath), user, pass))
	router.DELETE("/api/paths/*path", m.BasicAuth(m.AcceptJSON(controllers.DeletePath), user, pass))
	router.GET("/data/*path", controllers.GetPathData)
	router.GET("/help", controllers.Helppage)
	router.ServeFiles("/assets/*filepath", http.Dir("assets"))

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func dieHard(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Println(r.URL.Path, err)
	w.WriteHeader(http.StatusInternalServerError)
}
