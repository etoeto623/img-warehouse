package main

import (
	"flag"
	"fmt"

	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
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
	daemonMode := flag.Bool("d", false, "run as a daemon service")

	flag.Parse()

	config := common.InitConfig()
	if nil != serverPort && *serverPort > 0 {
		config.ServerPort = *serverPort
	}
	if nil != imageViewApi && len(*imageViewApi) > 0 {
		config.ImageViewApi = *imageViewApi
	}
	if len(config.AesKey) <= 0 {
		config.AesKey = common.AES_KEY
	}

	if nil != daemonMode && *daemonMode {
		mainthread.Init(func() {
			hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
			if err := hk.Register(); nil != err {
				fmt.Printf("hotkey: failed to register hotkey: %v\n", err)
				return
			}
			defer hk.Unregister()

			quitKey := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyQ)
			if err2 := quitKey.Register(); nil != err2 {
				fmt.Printf("hotkey: failed to register hotkey: %v\n", err2)
				return
			}
			defer quitKey.Unregister()

			for {
				select {
				case <-hk.Keyup():
					meta := common.ImageUploadMeta{}
					meta.FromClipboard = true
					doSendImage(&config, &meta)
				case <-quitKey.Keyup():
					return
				}
			}
		})
		return
	}

	if nil != serverMode && *serverMode {
		server.StartServer(&config)
		return
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

	doSendImage(&config, &meta)
}

// 上传图片数据
func doSendImage(config *common.EnvConfig, meta *common.ImageUploadMeta) {
	fname, err := client.UploadImage(config, meta)
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
