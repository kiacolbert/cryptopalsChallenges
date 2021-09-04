package main

import (
	"bytes"
	"regexp"
	"testing"
)

func TestGetHexStringToBaseSixtyFourString(t *testing.T) {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	base64String := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	want := regexp.MustCompile(`\b`+base64String+`\b`)
	msg := GetBase64StringFromHexString(hexString)

	if !want.MatchString(msg) {
		t.Fatalf(`GetBase64StringFromHexString("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d") = %q, want match for %#q`, msg, want)
	}

}

func TestGetXor(t *testing.T) {
	a := []byte{1,2}
	b := []byte{1,2}
	msg := GetXor(a,b)
	want := []byte{0,0}

	if !bytes.Equal(msg, want) {
		t.Fatalf(`GetXor([1,2],[1,2]) = %q, want match for %#q`,msg,want)
	}

	a = []byte{0,0}
	b = []byte{1,1}
	msg = GetXor(a,b)
	want = []byte{1,1}

	if !bytes.Equal(msg, want) {
		t.Fatalf(`GetXor([1,2],[0,1]) = %q, want match for %#q`,msg,want)
	}
}