package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// The hex string should produce the base64 string.
func main() {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	baseSixFourString := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	
	hexBytes, err := hex.DecodeString(hexString) // decode string into hex byte array

	if err != nil {
		fmt.Println(err)
	}

	sixFourString := base64.RawStdEncoding.EncodeToString(hexBytes) // encode hex bytes to base 64 string
	
	if sixFourString != baseSixFourString {
		fmt.Printf("incorrect")
	}
	fmt.Printf("correct: %v\n", sixFourString)

}
