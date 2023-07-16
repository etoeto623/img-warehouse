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

func GetRandomString(size int) string {

	var str []byte
	charPool := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	for len(str) < size {
		str = append(str, charPool[rand.Intn(len(charPool))])
	}
	return string(str)
}
