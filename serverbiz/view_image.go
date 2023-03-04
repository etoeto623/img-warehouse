package serverbiz

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"neolong.me/img-warehouse/common"
)

func getImageContent(path string) (*common.ImgViewRespData, error) {
	imgGetUrl := fmt.Sprintf("%s%s", common.ALIST_URL, common.IMAGE_GET_API)

	payload := strings.NewReader(genImgViewParamStr(path))

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

func genImgViewParamStr(path string) string {
	paramJson := common.ImgViewParam{
		Path:     path,
		Password: common.ALIST_PASSWORD,
	}

	bytes, _ := json.Marshal(paramJson)
	return string(bytes)
}
