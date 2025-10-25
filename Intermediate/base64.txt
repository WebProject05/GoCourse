// This is not an encryption method, can be decrypted easily.

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// Data transmitted should be encoded before transmitting
	// Base 64 is used in url's, text transmission etc.
	data := []byte("Hello base 64 encoding")
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(data)    // This is a byte array with ASCII values for each character
	fmt.Println(encoded) // This is the encoded string and with two "==" at the end as for padding

	// Decoding the encoded
	decoded, err := base64.StdEncoding.DecodeString(encoded) // Return values are byte slice and error
	if err != nil {
		fmt.Println("Error decoding:", err)
	}
	fmt.Println(decoded)
	fmt.Println(string(decoded))

	dataurl := []byte("https://udemy.com")
	// In URl encoding we need to avoind "/" and "+"
	urlEncoded := base64.URLEncoding.EncodeToString(dataurl)
	fmt.Println("URL encoded safely:", urlEncoded)

	decodedURL, err := base64.URLEncoding.DecodeString(urlEncoded)
	if err != nil {
		fmt.Println("Error decoding the url:", err)
	}
	fmt.Println(decodedURL)
	fmt.Println(string(decodedURL))
}
