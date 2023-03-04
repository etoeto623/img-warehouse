package server

import (
	"net/http"

	"neolong.me/neotools/cipher"
)

func StartServer() {
	http.HandleFunc("/view/", imgViewHandler)
	http.ListenAndServe(":6230", nil)
}

func imgViewHandler(w http.ResponseWriter, r *http.Request) {
	vid := r.URL.Query().Get("vid")

	cipher.AesDecryptString(vid)

	w.Write([]byte("hello"))
}
