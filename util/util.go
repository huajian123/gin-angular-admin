package util

import (
	"math/rand"
	"time"
)

// 生成一个随机的10位字符串
func RandomString(n int) string {
	var letters = []byte("asdfdsafasdfgdfsgfhgfdhadsflajfljsdklfjl")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	println(string(result))
	return (string(result))
}
