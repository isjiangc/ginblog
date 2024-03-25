package scrypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// ScryptPw 生成密码
func ScryptPw(password string) string {
	const cost = 10
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPw)
}
