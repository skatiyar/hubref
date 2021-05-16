package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/julienschmidt/httprouter"
	"github.com/skatiyar/hubref/controllers"
	"github.com/skatiyar/hubref/database"
	"github.com/skatiyar/hubref/middlewares"
)

var (
	port  = os.Getenv("PORT")
	user  = os.Getenv("USERNAME")
	pass  = os.Getenv("PASSWORD")
	dbURL = os.Getenv("DATABASE_URL")
)

func main() {
	if port == "" || user == "" || pass == "" {
		log.Fatal("PORT, USERNAME & PASSWORD must be set")
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
	router.GET("/", middlewares.BasicAuth(controllers.Homepage, user, pass))
	router.GET("/api/paths", middlewares.BasicAuth(controllers.GetPaths, user, pass))
	router.GET("/api/paths/*path", middlewares.BasicAuth(controllers.GetPath, user, pass))
	router.POST("/api/paths", middlewares.BasicAuth(controllers.CreatePath, user, pass))
	router.PUT("/api/paths/*path", middlewares.BasicAuth(controllers.EditPath, user, pass))
	router.DELETE("/api/paths/*path", middlewares.BasicAuth(controllers.DeletePath, user, pass))
	router.GET("/data/*path", controllers.GetPath)
	router.ServeFiles("/assets/*filepath", http.Dir("assets"))

	log.Fatal(http.ListenAndServe(":"+port, router))
}
