package tools

import (
	"gin-vue/dao"
	"math/rand"
	"time"
)

// RandomString 随机用户名产生函数
func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for key := range result {
		result[key] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func IsPhoneExist(phone string) bool {
	user := dao.Mgr.GetUserByPhone(phone)
	if user.ID != 0 {
		return true
	}
	return false
}
