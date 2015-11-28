package main

import (
	"github.com/goforgery/forgery2"
	"github.com/goforgery/responsetime"
)

func main() {
	app := f.CreateApp()
	app.Use(responsetime.Create())
	app.Listen(3000)
}
