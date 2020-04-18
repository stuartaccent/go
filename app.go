package main

import (
	"app/conf"
	"app/models"
	"app/routing"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	// wait for a db connection before we continue
	// useful when waiting for a docker container to become ready
	conf.WaitForDB()
}

func main() {
	// setup database
	db, _ := conf.ConnectDB()
	defer db.Close()
	models.SetDatabase(db)
	models.AutoMigrate()

	// setup validator
	conf.SetValidator()

	// setup global app context
	context := &conf.AppContext{
		Db:          db,
		CookieStore: conf.CookieStore,
		Templates:   conf.ParseTemplates("templates"),
	}

	// routing
	r := mux.NewRouter()
	// routes
	r = routing.BaseRoutes(r, context)
	r = routing.StaticRoutes(r, context)
	r = routing.AuthRoutes(r, context)
	r = routing.SampleLockedRoutes(r, context)

	http.Handle("/", r)

	// serve
	log.Fatal(http.ListenAndServe(":80", nil))
}
