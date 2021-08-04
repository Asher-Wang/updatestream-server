package main

import (
	// "encoding/json"
	"github.com/gin-gonic/gin"
	"log"

	passport "github.com/hotstar/passport-go"
)

var port = 3000

const (
	UserTokenHeader        = "X-HS-UserToken"
	IsUserTokenRefreshed   = "X-HS-IsUserTokenRefreshed"
	UserTokenAuthenticated = "X-HS-TS-Authenticated"
)

func main() {
	log.Printf("Auth service (HTTP) running on %d...\n", port)

	r := gin.Default()
	v1 := r.Group("/v1")
	v1.Use(passport.PassportMiddleware())
	{
		v1.GET("/user", func(c *gin.Context) {
			user := passport.GetUser(c)
			c.JSON(200, user)
		})
		v1.GET("/device", func(c *gin.Context) {
			device := passport.GetDevice(c)
			c.JSON(200, device)
		})
	}
	r.Run(":3000")
}
