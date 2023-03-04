package server

import (
	"net/http"

	"neolong.me/img-warehouse/common"
	"neolong.me/img-warehouse/serverbiz"
	"neolong.me/neotools/cipher"
)

func StartServer() {
	http.HandleFunc("/view/", imgViewHandler)
	http.ListenAndServe(":6230", nil)
}

func imgViewHandler(w http.ResponseWriter, r *http.Request) {
	vid := r.URL.Query().Get("vid")

	path, err := cipher.AesDecryptString(vid, common.AES_KEY)
	if nil != err {
		w.Write(common.StrToBytes("invalid image"))
		return
	}

	imgData, err := serverbiz.GetImageData(path)
	if nil != err {
		w.Write(common.StrToBytes("fail to get image"))
		return
	}

	http.Redirect(w, r, imgData.RawUrl, http.StatusFound)
}
