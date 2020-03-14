package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	keyEncoded := []byte("vVibSI4hUFkLIsuW5SfGnFLXAXnEcwzTS472r6D+x3Y=")
	nonceEncoded := []byte("rMr7iSz1k0oa1TrU")

	//Method-1
	keyDecoded := make([]byte, base64.StdEncoding.DecodedLen(len(keyEncoded)))
	l1, _ := base64.StdEncoding.Decode(keyDecoded, keyEncoded)

	//Method-2
	nonceDecoded, _ := base64.StdEncoding.DecodeString(string(nonceEncoded))
	//textDecoded, _ := base64.StdEncoding.DecodeString(encText)

	//key1, _ := GetBytes(32)
	//nonce, _ = GetBytes(12)

	fmt.Println("Method-1")
	fmt.Printf("\n\n****Encoded Values****\n\n")

	fmt.Println("Encoded string (default) ", keyEncoded)
	fmt.Printf("Encoded string (string) %s\n", keyEncoded)
	fmt.Println("Encoded string length ", len(keyEncoded))

	fmt.Printf("\n\n****Decoded Values****\n\n")

	fmt.Println("Decoded string (default) ", keyDecoded[:l1])
	fmt.Printf("Decoded string (string) %s\n", keyDecoded[:l1])
	fmt.Printf("Decoded string (hex) % x\n", keyDecoded[:l1])
	fmt.Printf("Decoded string(double quoted/escaped) %+q\n", keyDecoded[:l1])
	fmt.Println("Decoded string length ", len(keyDecoded[:l1]))

	fmt.Println("Method-2")
	fmt.Printf("\n\n****Encoded Values****\n\n")

	fmt.Println("Encoded string (default) ", nonceEncoded)
	fmt.Printf("Encoded string (string) %s\n", nonceEncoded)
	fmt.Println("Encoded string length ", len(nonceEncoded))

	fmt.Printf("\n\n****Decoded Values****\n\n")

	fmt.Println("Decoded string (default) ", nonceDecoded)
	fmt.Printf("Decoded string (string) %s\n", nonceDecoded)
	fmt.Printf("Decoded string (hex) % x\n", nonceDecoded)
	fmt.Printf("Decoded string(double quoted/escaped) %+q\n", nonceDecoded)
	fmt.Println("Decoded string length ", len(nonceDecoded))

}
