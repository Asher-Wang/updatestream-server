package main

import (
	// "encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hotstar/passport-go"
)

var port = 3000

const (
	UserTokenHeader        = "X-HS-UserToken"
	IsUserTokenRefreshed   = "X-HS-IsUserTokenRefreshed"
	UserTokenAuthenticated = "X-HS-TS-Authenticated"
	PassportHeader         = "X-HS-Passport"
)

func main() {
	log.Printf("Auth service (HTTP) running on  %d...\n", port)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/user", func(c *gin.Context) {
			user := passport.GinGetUser(c)
			c.JSON(200, gin.H{"user": user, "request_headers": c.Request.Header})
		})
		v1.GET("/device", func(c *gin.Context) {
			device := passport.GinGetDevice(c)
			c.JSON(200, gin.H{"device": device, "request_headers": c.Request.Header})
		})
		v1.GET("/passport", func(c *gin.Context) {
			p := passport.GinGetPassport(c)
			c.JSON(200, gin.H{"passport": p, "request_headers": c.Request.Header})
		})
	}
	r.Run(":8080")
}
