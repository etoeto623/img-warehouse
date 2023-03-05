package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"neolong.me/img-warehouse/common"
	"neolong.me/img-warehouse/serverbiz"
	"neolong.me/neotools/cipher"
)

func StartServer(config *common.EnvConfig) {
	port := strconv.Itoa(config.ServerPort)
	fmt.Println(":" + port + "/" + config.ImageViewApi)
	http.HandleFunc("/"+config.ImageViewApi, func(w http.ResponseWriter, r *http.Request) {
		imgViewHandler(w, r, config)
	})
	http.ListenAndServe(":"+port, nil)
}

func imgViewHandler(w http.ResponseWriter, r *http.Request, config *common.EnvConfig) {
	vid := r.URL.Query().Get("vid")

	path, err := cipher.AesDecryptString(vid, config.AesKey)
	if nil != err {
		w.Write(common.StrToBytes("invalid image"))
		return
	}

	imgData, err := serverbiz.GetImageData(path, config)
	if nil != err {
		w.Write(common.StrToBytes("fail to get image"))
		return
	}

	content, err := readContentFromRawUrl(imgData.RawUrl)
	if nil != err {
		w.Write(common.StrToBytes("read image data fila"))
		return
	}

	w.Write(content)
}

func readContentFromRawUrl(rawUrl string) ([]byte, error) {
	resp, err := http.DefaultClient.Get(rawUrl)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
