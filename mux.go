package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()

	// Redirect all request to "mux.ro" if not already here
	e.Use(func(c *echo.Context) {
		if c.Request.Host != "mux.ro" {
			c.Redirect(302, "http://mux.ro")
		}
	})

	// Serve the stuff
	e.Static("", "public")

	e.Run(":8080")
}
