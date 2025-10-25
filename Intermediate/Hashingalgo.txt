package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

// SHA 256 - 256 bit or 32 byte hash
// SHA 512 - 512 bit or 54 byte hash
// Same input will produce the same hash output everytime
// The hashed key is not reversiable

func main() {
	password := "password123"
	// hashedPassword := sha256.Sum256([]byte(password))
	// fmt.Println(password)
	// fmt.Println(hashedPassword)	// This is a byte slice

	// fmt.Printf("SHA 256 hash hex value: %x \n", hashedPassword)

	// password_company := "sano553"
	// hashedPassword1 := sha512.Sum512([]byte(password_company))
	// fmt.Printf("SHA 512 hash hex value: %x \n", hashedPassword1)

	// Salting: it is a value which is added to the information before hasing
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt:", err)
	}

	// This is encodedSalt, so that we can store this in the database along with the password hashed and then compare them while login
	encodedSalt := base64.StdEncoding.EncodeToString(salt)

	fmt.Println("Encoded salt:", encodedSalt)

	SignUphash := hashPassword(password, salt)
	fmt.Println("The hashed password with salt in it is:", SignUphash)

	// LOGIN SECTION
	// Veryify the password
	inputPass := "password12"
	decodedSalt, err1 := base64.StdEncoding.DecodeString(encodedSalt)
	if err1 != nil {
		fmt.Println("Error:", err)
		return
	}
	
	loginHash := hashPassword(inputPass, decodedSalt)
	if loginHash == SignUphash {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Wrong Password. Try again")
	}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16) // byte slice of 16 bytes
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Function to hash password
func hashPassword(password string, salt []byte) string {
	saltedPass := append([]byte(password), salt...)

	hash := sha256.Sum256(saltedPass)
	return hex.EncodeToString(hash[:])
}
