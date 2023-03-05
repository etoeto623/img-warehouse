package serverbiz

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"neolong.me/img-warehouse/common"
)

func GetImageData(path string, cfg *common.EnvConfig) (*common.ImgViewRespData, error) {
	imgGetUrl := fmt.Sprintf("%s%s", cfg.AlistUrl, common.IMAGE_GET_API)

	payload := strings.NewReader(genImgViewParamStr(path, cfg))

	resp, err := http.Post(imgGetUrl, "application/json;charset=UTF-8", payload)
	if nil != err {
		return nil, err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, err
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

	bytes, _ := json.Marshal(paramJson)
	return string(bytes)
}
