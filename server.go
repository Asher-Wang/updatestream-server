package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var port = 3000

const (
	UserTokenHeader        = "X-HS-UserToken"
	IsUserTokenRefreshed   = "X-HS-IsUserTokenRefreshed"
	UserTokenAuthenticated = "X-HS-TS-Authenticated"
)

type Handler struct {
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	authenticated := req.Header.Get(strings.ToLower(UserTokenAuthenticated))
	if authenticated != "true" {
		rErr(resp, 403, "unauthenticated")

		return
	}
	isTokenRefreshed := req.Header.Get(strings.ToLower(IsUserTokenRefreshed))
	if isTokenRefreshed == "true" {
		refreshedToken := req.Header.Get(strings.ToLower(UserTokenHeader))
		if refreshedToken == "" {
			rErr(resp, 403, "Empty passport")
			return
		}
		resp.Header().Add(UserTokenHeader, refreshedToken)
	}

	resp.Header().Add("Content-Type", "application/json")
}

func main() {
	log.Printf("Auth service (HTTP) running on %d...\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), &Handler{})
	if err != nil {
		panic(err)
	}
}

func rErr(resp http.ResponseWriter, statusCode int, message string) {
	resp.WriteHeader(statusCode)
	resp.Write([]byte(message))

	log.Printf("[Response] %d: %s\n", statusCode, message)
}
