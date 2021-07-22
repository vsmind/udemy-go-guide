package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	response, error := http.Get("http://google.com")

	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}
	// Create a byte slice with 99999 empty elements
	byteSlice := make([]byte, 99999)
	response.Body.Read(byteSlice)
	fmt.Println(string(byteSlice))

	io.Copy(os.Stdout, response.Body)
	fmt.Println("_________________________________")
	lw := logWriter{}
	io.Copy(lw, response.Body)

}

func (logWriter) Write(bs []byte) (size int, err error) {
	fmt.Println(string(bs))
	wBytes := len(bs)
	fmt.Println("Number of written bytes: ", wBytes)

	return wBytes, nil
}
