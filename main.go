package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/vichar/fizzbuzz"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/fizzbuzz/:number", func(c echo.Context) error {
		number := c.Param("number")
		n, error := strconv.Atoi(number)
		if error != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": error.Error(),
			})
		}
		if n > 5 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": number + " is not a supported number.",
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": fizzbuzz.Say(n),
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
