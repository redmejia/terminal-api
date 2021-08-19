package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("name ", os.Getenv("NAME"))
}
