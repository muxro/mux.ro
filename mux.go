package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()

	// Serve the stuff
	e.Static("", "public")

	e.Run(":8080")
}
