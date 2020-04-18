package conf

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
)

type AppContext struct {
	Db          *gorm.DB
	CookieStore *sessions.CookieStore
	Templates   *template.Template
}

func (ac AppContext) TemplateResponse(w http.ResponseWriter, tmpl string, data interface{}) (int, error) {
	if err := ac.Templates.ExecuteTemplate(w, tmpl, data); err != nil {
		log.Fatal(err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

type AppHandler struct {
	*AppContext
	H func(*AppContext, http.ResponseWriter, *http.Request) (int, error)
}

func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := ah.H(ah.AppContext, w, r)
	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}
