package main

import (
	"github.com/magmel48/go-projects/networking/portschecker"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	portschecker.CheckPorts()
}
