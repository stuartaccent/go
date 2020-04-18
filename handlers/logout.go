package handlers

import (
	"app/conf"
	"net/http"
)

func Logout(appContext *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	session, _ := appContext.CookieStore.Get(r, "session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
	return http.StatusFound, nil
}
