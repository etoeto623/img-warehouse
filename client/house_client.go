package client

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

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
	refreshWarehouseList(cfg)

	encryptedFname, _ := cipher.AesEncryptString("/"+fname, cfg.AesKey)
	return encryptedFname, nil
}

// 刷新仓库的资源列表
func refreshWarehouseList(cfg *common.EnvConfig) {
	data := common.ImageListMeta{
		Page:     1,
		PerPage:  0,
		Path:     "/",
		Password: cfg.AlistPassword,
		Refresh:  true,
	}
	refreshUrl := fmt.Sprintf("%s%s", cfg.AlistUrl, common.IMAGE_LIST_API)
	payload := strings.NewReader(common.StructToString(data))
	if req, err := http.NewRequest(http.MethodPost, refreshUrl, payload); nil == err {
		req.Header.Add("Content-Type", "application/json; charset=utf-8")
		http.DefaultClient.Do(req)
	}
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

// 发送数据到alist服务端
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
	if len(cfg.AlistPassword) <= 0 {
		// 没指定密码，则从服务器获取
		req.Header.Add("Authorization", acquireToken(cfg))
	} else {
		// 默认使用guest用户
		req.Header.Add("Password", cfg.AlistPassword)
	}

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

func acquireToken(cfg *common.EnvConfig) string {
	tokenUrl := fmt.Sprintf("%s%s", cfg.AlistUrl, common.SERVER_TOKEN_API)
	req, err := http.NewRequest(http.MethodPut, tokenUrl, nil)
	if nil != err {
		fmt.Println("shit happens", err.Error())
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		fmt.Println("shit happens", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Println("shit happens", err.Error())
		os.Exit(1)
	}

	tokenBytes, err := cipher.AesDecrypt(respBytes, cfg.AesKey)
	if nil != err {
		fmt.Println("shit happens", err.Error())
		os.Exit(1)
	}

	return string(tokenBytes)
}
