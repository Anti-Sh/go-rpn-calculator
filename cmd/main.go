package main

import (
	"fmt"
	"github.com/Anti-Sh/go-rpn-calculator/internal/application"
)

func main() {
	app := application.NewApplication()

	err := app.RunServer()
	if err != nil {
		fmt.Println("Server error", err)
		return
	}
}
