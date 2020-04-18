package handlers

import "net/http"

func ServiceWorker(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/js/service-worker.js")
}
