package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// md加密
func md5Hash(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	original := "admin"
	firstHash := md5Hash(original)
	secondHash := md5Hash(firstHash)

	fmt.Println("First MD5 hash:", firstHash)
	fmt.Println("Second MD5 hash:", secondHash)
}
