package common

import (
	"encoding/json"
	"math/rand"
	"time"
)

func StrToBytes(data string) []byte {
	return []byte(data)
}

func GenTimestamp() string {
	return time.Now().Format("060102150405")
}

func StructToString(data interface{}) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

func BytesToMap(data []byte) (map[string]interface{}, error) {
	dataMap := make(map[string]interface{})
	err := json.Unmarshal(data, &dataMap)
	return dataMap, err
}

func GetRandomString(size int) string {

	var str []byte
	charPool := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	for len(str) < size {
		str = append(str, charPool[rand.Intn(len(charPool))])
	}
	return string(str)
}

func RespSuccess(respMap map[string]interface{}) bool {
	return int(respMap["code"].(float64)) == 200
}
