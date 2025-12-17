package main

import (
	"fmt"
	"go-fiber/utils"
)

func main() {
	hash, _ := utils.HashPassword("12345")
	fmt.Println(hash)
}