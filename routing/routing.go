package routing

import (
	"app/conf"
	"app/handlers"
	"app/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func BaseRoutes(r *mux.Router, ctx *conf.AppContext) *mux.Router {
	r.Handle("/", middleware.CurrentUser(conf.AppHandler{ctx, handlers.Index})).Methods("GET")
	return r
}

func AuthRoutes(r *mux.Router, ctx *conf.AppContext) *mux.Router {
	s := r.PathPrefix("/auth").Subrouter()
	s.Handle("/login", conf.AppHandler{ctx, handlers.LoginGet}).Methods("GET")
	s.Handle("/login", conf.AppHandler{ctx, handlers.LoginPost}).Methods("POST")
	s.Handle("/logout", conf.AppHandler{ctx, handlers.Logout}).Methods("GET")
	s.Handle("/password/reset", conf.AppHandler{ctx, handlers.PWResetGet}).Methods("GET")
	s.Handle("/password/reset", conf.AppHandler{ctx, handlers.PWResetPost}).Methods("POST")
	s.Handle("/password/reset/done", conf.AppHandler{ctx, handlers.PWResetDone}).Methods("GET")
	s.Handle("/password/reset/confirm", conf.AppHandler{ctx, handlers.PWResetConfirmGet}).Methods("GET")
	s.Handle("/password/reset/confirm", conf.AppHandler{ctx, handlers.PWResetConfirmPost}).Methods("POST")
	s.Handle("/password/reset/complete", conf.AppHandler{ctx, handlers.PWResetComplete}).Methods("GET")
	return r
}

func StaticRoutes(r *mux.Router, ctx *conf.AppContext) *mux.Router {
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("/service-worker.js", handlers.ServiceWorker).Methods("GET")
	return r
}

func SampleLockedRoutes(r *mux.Router, ctx *conf.AppContext) *mux.Router {
	s := r.PathPrefix("/").Subrouter()
	s.Use(middleware.CurrentUser, middleware.AuthRequired)
	s.Handle("/restricted", conf.AppHandler{ctx, handlers.Index}).Methods("GET")
	return r
}
