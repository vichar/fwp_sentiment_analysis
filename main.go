package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/vichar/rssfeed"
)

func main() {

	channels := make([]rssfeed.RSSChannel, 0)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.POST("/programs", func(c echo.Context) error {
		feed := c.FormValue("rssURL")
		response, error := rssfeed.HTTPGet(feed)
		if error != nil {
			e.Logger.Error("HTTP Client Error")
			return c.JSON(http.StatusNotAcceptable, "Error Parsing HTTP Response")
		}
		httpResponse := rssfeed.ParseHTTPResponse(response)
		channel, parseError := rssfeed.ParseRSSData(httpResponse.Body)
		if parseError != nil {
			e.Logger.Error("Error Parsing HTTP Response")
			return c.JSON(http.StatusNotAcceptable, "Error Parsing HTTP Response")
		}
		channels = append(channels, channel)
		channelsJSON, jsonError := json.Marshal(channels)
		if jsonError != nil {
			e.Logger.Error("Error Parsing Data")
			return c.JSON(http.StatusNotAcceptable, "Error Parsing Data")
		}

		return c.JSON(http.StatusOK, json.RawMessage(string(channelsJSON)))

	})

	e.GET("/programs", func(c echo.Context) error {
		channelsJSON, jsonError := json.Marshal(channels)
		if jsonError != nil {
			e.Logger.Error("Error Parsing Data")
			return c.JSON(http.StatusNotAcceptable, "Error Parsing Data")
		}
		return c.JSON(http.StatusOK, json.RawMessage(string(channelsJSON)))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
