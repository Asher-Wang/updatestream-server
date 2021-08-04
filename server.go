package main

import (
	"encoding/json"
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
			userb, err := json.Marshal(user)
			if err != nil {
				c.AbortWithError(400, err)
			}
			c.JSON(200, gin.H{
				"user": string(userb),
			})
		})
		v1.GET("/device", func(c *gin.Context) {
			device := passport.GetDevice(c)
			deviceb, err := json.Marshal(device)
			if err != nil {
				c.AbortWithError(400, err)
			}
			c.JSON(200, gin.H{
				"user": string(deviceb),
			})
		})
	}
	r.Run()
}
