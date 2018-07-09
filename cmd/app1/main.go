package main

import (
	"fmt"

	"github.com/philoticnet/multiple_apps/cmd/app1/utils"
)

func main() {
	fmt.Printf("start app1")
	utils.WriteFile("app1.txt", "was written by app1")
}
