// Copyright (c) 2022 arcezd
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		name := c.DefaultQuery("name", "World")
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello %s", name),
		})
	})

	router.GET("/os", func(c *gin.Context) {
		c.String(200, runtime.GOOS)
	})
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))
}
