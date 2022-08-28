package main

import (
	"github.com/gin-gonic/gin"
  "net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../public/views/*.html")
	r.Static("../public", "../public/")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})


	r.Run("127.0.0.1:3000") // listen and serve on 0.0.0.0:8080
}