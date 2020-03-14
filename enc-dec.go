package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {

	textBytes := []byte("Service is worship")

	key, _ := GetBytes(32)
	nonce, _ := GetBytes(12)

	cipherTextBytes, err := EncryptText(textBytes, key, nonce)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ciphertext: %s\n", cipherTextBytes)

	plainTextBytes, err := DecryptText(cipherTextBytes, key, nonce)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("plaintext: %s\n", plainTextBytes)

}

// Get byte array with random size
func GetBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return nil, err
	}
	return bytes, nil
}

//Decrypt CipherText
func DecryptText(data []byte, symmKey []byte, nonce []byte) ([]byte, error) {

	block, err := aes.NewCipher(symmKey)
	if err != nil {
		fmt.Println("err with cipher ", err)
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("err with GCM ", err)
		return nil, err
	}

	plainTextBytes, err := aesgcm.Open(nil, nonce, data, nil)
	if err != nil {
		fmt.Println("err with opening ", err)
		return nil, err
	}
	return plainTextBytes, nil
}

//Encrypt the given Plain text
func EncryptText(data []byte, symmKey []byte, nonce []byte) ([]byte, error) {

	block, err := aes.NewCipher(symmKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	cipherTextBytes := aesgcm.Seal(nil, nonce, data, nil)
	if err != nil {
		fmt.Println("err with sealing ", err)
		return nil, err
	}
	return cipherTextBytes, nil
}
