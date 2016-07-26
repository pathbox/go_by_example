package main

import (
	// "fmt"
	"net/http"
)

func main() {
	http.Handle("/test", http.FileServer(http.Dir("./")))
	http.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir("./"))))
	http.Handle("/tmpfile", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))
	http.ListenAndServe(":9999", nil)

}
