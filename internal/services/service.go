package services

import (
	"fmt"
	"net/http"
)

func Greet(w http.ResponseWriter, req *http.Request) {
	nameStr := req.PathValue("name")
	fmt.Fprintf(w, "Hello, %s :)\n", nameStr)
}
