package common

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func InitConfig() EnvConfig {
	configDir := os.Getenv("HOME")
	if len(configDir) <= 0 {
		configDir = "./"
	}

	cfg := EnvConfig{}
	bytes, err := ioutil.ReadFile(configDir + "/.imghouse.cfg")
	if nil != err {
		return cfg
	}

	json.Unmarshal(bytes, &cfg)
	return cfg
}
