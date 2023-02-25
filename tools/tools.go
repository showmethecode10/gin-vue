package tools

import (
	"math/rand"
	"time"
)

//随机用户名产生函数
func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for key := range result {
		result[key] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
