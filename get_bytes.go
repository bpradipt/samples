package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
)

func main() {
	size := flag.Int("size", 32, "Required byte size[default=32")

	flag.Parse()

	key, err := GetBytes(*size)
	if err != nil {
		fmt.Printf("Error in generating random bytes %s", err)
	}
	keyEncoded := base64.StdEncoding.EncodeToString(key)
	fmt.Printf("Bytes[%d] = %s\n", *size, key)
	fmt.Printf("Base64 Encoded key %s\n", keyEncoded)

}

//Return byte array of requested size
func GetBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return nil, err
	}
	return bytes, nil
}
