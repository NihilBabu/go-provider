package handlers

import "net/http"

//IndexHandler is ...
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hlo"))
}
