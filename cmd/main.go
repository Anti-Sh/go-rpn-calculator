package main

import "github.com/Anti-Sh/go-rpn-calculator/internal/application"

func main() {
	app := application.NewApplication()

	app.RunServer()
}
