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
	PassportHeader         = "X-HS-Passport"
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
		v1.GET("/passport", func(c *gin.Context) {
			for k, vals := range c.Request.Header {
				log.Printf("%s", k)
				for _, v := range vals {
					log.Printf("\t%s", v)
				}
			}
			passportHeader := c.GetHeader(PassportHeader)
			pbPassport, err := passport.DecodePassport(passportHeader)
			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
			}
			c.JSON(200, pbPassport)
		})
	}
	r.Run(":8080")
}
