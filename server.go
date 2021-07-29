package main

import (
	"fmt"
	"log"
	"net/http"
)

var port = 3000

type Handler struct {
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	authenticated := req.Header.Get("X-HS-TS-Authenticated")
	if authenticated != "true" {
		rErr(resp, 403, "unauthenticated")
		return
	}
	isTokenRefreshed := req.Header.Get("X-HS-IsUserTokenRefreshed")
	if isTokenRefreshed == "true" {
		refreshedToken := req.Header.Get("X-HS-UserToken")
		if refreshedToken == "" {
			rErr(resp, 403, "Empty passport")
			return
		}
		resp.Header().Add("X-HS-UserToken", refreshedToken)
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
