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

	sixFourString := GetBase64StringFromHexString(hexString)
	if sixFourString != baseSixFourString {
		fmt.Printf("incorrect")
	}
	fmt.Printf("correct: %v\n", sixFourString)

	ans := getXor(hexBytes, hexBytes)
	fmt.Println("ans", ans)

func GetBase64StringFromHexString(hexString string) string{
	hexBytes, err := hex.DecodeString(hexString) // decode string into hex byte array
	if err != nil {
		fmt.Println(err)
	}
	return base64.RawStdEncoding.EncodeToString(hexBytes) // encode hex bytes to base 64 string
}

// Write a function that takes two equal-length buffers and produces their XOR combination.
func getXor(a []byte, b []byte) []byte {
	if len(a) != len(b) {
		return nil;
	}
	res := make([]byte, len(a))
	for i := range(a) {
		res[i] = a[i] ^ b[i]
	}
	return res
}
