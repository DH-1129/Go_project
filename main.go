package main

import (
	"fmt"

	"dhui.com/routes"
)

func main() {
	r := routes.SetupRoute()
	if err := r.Run(":8000"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
