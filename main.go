package main

import (
	"fmt"

	"github.com/alextanhongpin/rest-api-101/config"
)

func main() {
	config := config.Read()
	fmt.Println(config.Port)
}
