package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var jsonDataString string

	r.GET("/brick", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://cdn.onebrick.io/sandbox-widget/v1/?accessToken=public-sandbox-4bec8253-2ee3-478c-a17d-9168715cd793&redirect_url=https://6ba4-180-252-172-19.ap.ngrok.io")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, jsonDataString)
	})

	r.POST("/", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		jsonDataString = string(jsonData)

		c.String(http.StatusOK, "https://6ba4-180-252-172-19.ap.ngrok.io")
	})
	r.Run("localhost:3000")
}
