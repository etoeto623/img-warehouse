package main

import (
	"flag"
	"fmt"

	"golang.design/x/clipboard"
	"neolong.me/img-warehouse/client"
	"neolong.me/img-warehouse/common"
	"neolong.me/img-warehouse/server"
)

func main() {

	serverMode := flag.Bool("s", false, "run as server")
	// clientMode := flag.Bool("c", true, "run as client")

	fileName := flag.String("fn", "", "specify image file name")
	localFile := flag.String("f", "", "local image file path")
	serverPort := flag.Int("p", 0, "server port")
	imageViewApi := flag.String("api", "", "image view api prefix in server")

	flag.Parse()

	config := common.InitConfig()
	if nil != serverPort && *serverPort > 0 {
		config.ServerPort = *serverPort
	}
	if nil != imageViewApi && len(*imageViewApi) > 0 {
		config.ImageViewApi = *imageViewApi
	}

	if nil != serverMode && *serverMode {
		server.StartServer(&config)
		return
	}
	if len(config.AesKey) <= 0 {
		config.AesKey = common.AES_KEY
	}

	// client mode
	meta := common.ImageUploadMeta{}
	if nil != localFile {
		meta.LocalFilePath = *localFile
	}
	if nil != fileName {
		meta.FileName = *fileName
	}
	meta.FromClipboard = nil == localFile || len(*localFile) <= 0

	fname, err := client.UploadImage(&config, &meta)
	if nil != err {
		fmt.Println("error:", err.Error())
	}
	imgUrl := fmt.Sprintf("%s%s?vid=%s", config.ServerUrl, config.ImageViewApi, fname)
	// write uploaded image url to clipboard
	if nil == clipboard.Init() {
		clipboard.Write(clipboard.FmtText, common.StrToBytes(imgUrl))
	}
	fmt.Println(imgUrl)
}
