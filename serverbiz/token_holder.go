package serverbiz

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"neolong.me/img-warehouse/common"
	"neolong.me/neotools/cipher"
)

var AlistToken string

func FetchToken(cfg *common.EnvConfig, times int) string {
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
	return FetchToken(cfg, times+1)
}

func RefreshToken(cfg *common.EnvConfig, times int) string {
	AlistToken = ""
	return FetchToken(cfg, times)
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
