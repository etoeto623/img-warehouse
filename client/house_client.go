package client

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"golang.design/x/clipboard"
	"neolong.me/img-warehouse/common"
	"neolong.me/neotools/cipher"
)

func UploadImage(cfg *common.EnvConfig, meta *common.ImageUploadMeta) (string, error) {
	imgData, err := getImageContentFromMeta(meta)
	if nil != err {
		return "", err
	}
	if nil == imgData || len(imgData) <= 0 {
		return "", errors.New("no iamge data in clipboard")
	}

	fname, err := sendImgDataToAlist(imgData, cfg, meta)
	if nil != err {
		return "", err
	}

	encryptedFname, _ := cipher.AesEncryptString("/"+fname, cfg.AesKey)
	return encryptedFname, nil
}

func getImageContentFromMeta(meta *common.ImageUploadMeta) ([]byte, error) {
	if meta.FromClipboard {
		return getImgDataFromClipboard()
	}

	return ioutil.ReadFile(meta.LocalFilePath)
}

func getImgDataFromClipboard() ([]byte, error) {
	if err := clipboard.Init(); nil != err {
		return nil, err
	}

	return clipboard.Read(clipboard.FmtImage), nil
}

func sendImgDataToAlist(data []byte, cfg *common.EnvConfig, meta *common.ImageUploadMeta) (string, error) {

	uploadUrl := fmt.Sprintf("%s%s", cfg.AlistUrl, common.IMAGE_UPLOAD_API)
	payload := bytes.NewReader(data)
	req, err := http.NewRequest(http.MethodPut, uploadUrl, payload)
	if nil != err {
		return "", err
	}

	fileName := common.GenTimestamp() + "_" + common.GetRandomString(6) + ".png"
	if meta.FromClipboard {
		fileName = "SS_" + fileName
	}

	req.Header.Add("Content-Length", strconv.Itoa(len(data)))
	req.Header.Add("Content-Type", "image/png")
	req.Header.Add("File-Path", "%2F"+fileName)
	req.Header.Add("Password", cfg.AlistPassword)

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		return "", err
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if nil != err {
		return "", err
	}
	return fileName, nil
}
