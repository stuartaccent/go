package handlers

import (
	"app/conf"
	"app/middleware"
	"app/models"
	"net/http"
)

func Index(appContext *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	context := struct {
		User *models.User
	}{
		User: middleware.GetUser(r.Context()),
	}
	return appContext.TemplateResponse(w, "indexHTML", context)
}
