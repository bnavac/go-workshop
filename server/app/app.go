package test

import (
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func helper(password string) []byte {
	newPass := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(newPass, 0)
	if err != nil {
		fmt.Printf("failed to hash")
		return nil
	}
	return hash
}
func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	hash := helper("password")
	fmt.Fprintf(w, "The hash of password is %q", hash)
}
func Headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
