package serverbiz

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"neolong.me/img-warehouse/common"
)

func GetImageData(path string, cfg *common.EnvConfig, couldRetry bool) (*common.ImgViewRespData, error) {
	imgGetUrl := fmt.Sprintf("%s%s", cfg.AlistUrl, common.IMAGE_GET_API)

	payload := strings.NewReader(genImgViewParamStr(path, cfg))

	req, err := http.NewRequest(http.MethodPost, imgGetUrl, payload)
	if nil != err {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	if len(cfg.AlistUser) > 0 {
		req.Header.Add("Authorization", FetchToken(cfg, 0))
	}
	resp, err := http.DefaultClient.Do(req)

	// resp, err := http.Post(imgGetUrl, "application/json;charset=UTF-8", payload)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}
	respMap, err := common.BytesToMap(bytes)
	if nil != err {
		return nil, err
	}
	if !common.RespSuccess(respMap) {
		if couldRetry {
			AlistToken = ""
			return GetImageData(path, cfg, false)
		}
		return nil, errors.New("img raw get error")
	}

	respJson := common.ImgViewResp{}
	err = json.Unmarshal(bytes, &respJson)
	if nil != err {
		return nil, err
	}
	return &respJson.Data, nil
}

func genImgViewParamStr(path string, cfg *common.EnvConfig) string {
	paramJson := common.ImgViewParam{
		Path:     path,
		Password: cfg.AlistPassword,
	}

	return common.StructToString(paramJson)
}
