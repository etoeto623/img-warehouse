package client

import (
	"fmt"

	"golang.design/x/hotkey"
	"golang.design/x/mainthread"
	"neolong.me/img-warehouse/common"
)

func InitHotkey(cfg *common.EnvConfig) {
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
				DoSendImage(cfg, &meta)
			case <-quitKey.Keyup():
				return
			}
		}
	})
}
