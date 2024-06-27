package main

import (
	"fmt"

	"github.com/laurati/api_product_user/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config.DBDriver)
}
