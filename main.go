package main

import (
	"github.com/magmel48/go-projects/security/caesar_cipher"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	caesar_cipher.CaesarCipher()
}
