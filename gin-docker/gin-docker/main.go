package main

import (
	    "net/http"
	    "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, map[string]interface{}{"hello":name})
	})
	router.Run(":8090")
}
