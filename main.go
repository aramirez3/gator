package main

import (
	"github.com/aramirez3/gator/internal/config"
)

func main() {
	config.Read()
	config.SetUser("angel")
	config.Read()
}
