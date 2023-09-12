package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"neolong.me/img-warehouse/common"
	"neolong.me/img-warehouse/serverbiz"
	"neolong.me/neotools/cipher"
)

var AlistToken string

func StartServer(config *common.EnvConfig) {
	port := strconv.Itoa(config.ServerPort)
	fmt.Println(":" + port + "/" + config.ImageViewApi)
	http.HandleFunc("/"+config.ImageViewApi, func(w http.ResponseWriter, r *http.Request) {
		imgViewHandler(w, r, config)
	})
	http.HandleFunc(common.SERVER_TOKEN_API, func(w http.ResponseWriter, r *http.Request) {
		// 返回token
		tk := fetchToken(config, 0)
		w.Write([]byte(tk))
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

func fetchToken(cfg *common.EnvConfig, times int) string {
	if times > 3 {
		return ""
	}
	if len(AlistToken) <= 0 {
		doLogin(cfg)
	}
	if len(AlistToken) > 0 {
		tk, _ := cipher.AesEncryptString(AlistToken, cfg.AesKey)
		return tk
	}
	return fetchToken(cfg, times+1)
}

// 进行登录
func doLogin(cfg *common.EnvConfig) error {
	loginUrl := fmt.Sprintf("%s%s", cfg.AlistUrl, common.LOGIN_API)
	payloadMap := make(map[string]string)
	payloadMap["password"] = cfg.AlistPassword
	payloadMap["username"] = cfg.AlistUser
	payloadMap["otp_code"] = ""
	payloadBytes, _ := json.Marshal(payloadMap)
	payload := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest(http.MethodPut, loginUrl, payload)
	if nil != err {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		return err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return err
	}
	respMap := make(map[string]interface{})
	if err = json.Unmarshal(respBytes, &respMap); nil != err {
		return err
	}
	// 判断是否登录成功
	respCode := int(respMap["code"].(float64))
	if respCode != 200 {
		return errors.New("login failed")
	}
	dataMap := respMap["data"].(map[string]interface{})
	AlistToken = dataMap["token"].(string)

	return nil
}
