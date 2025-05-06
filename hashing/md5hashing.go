package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func md5hashing(s string) string {
	var hash = md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func main() {
	var password string

	fmt.Scanln(&password)

	fmt.Println("the value fo the password is: ", md5hashing(password))
}
