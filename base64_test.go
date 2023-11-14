package base64

import (
	"fmt"
	"testing"
)

func TestEncodeText(t *testing.T) {
	fmt.Println(EncodeText("CYJ")) // Q1lK
}

func TestDecodeText(t *testing.T) {
	fmt.Println(DecodeText("Q1lK")) // CYJ
}
