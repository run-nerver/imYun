package libs

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// 加密密码
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, 10)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// 验证密码
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}