package base64

import (
	"fmt"
	"math"
)

func EncodeText(text string) string {

	var textBytes []byte = []byte(text)

	bits := binaryBitCutting(textBytes)
	resultBase64 := bitToBase64Encoding(bits)

	return resultBase64
}

func DecodeText(base64Text string) string {

	paddingCnt := 0
	for i := len(base64Text); i > 0; i-- {
		if base64Text[i-1] == '=' {
			paddingCnt++
		}
	}

	allBits := make([]byte, 0)
	for i := 0; i < len(base64Text); i++ {
		var result [6]byte = base64CharToSixBit(base64Text[i])
		allBits = append(allBits, result[:]...)
	}

	var resultText string = ""
	fmt.Printf("%b\n", allBits)

	loopLen := len(allBits) / 8
	for i := 0; i < loopLen; i++ {
		var curBit byte = 0
		for j := 0; j < 8; j++ {
			curBit |= allBits[i*8+j] << (7 - j)
		}
		if curBit == 0 && i+paddingCnt >= loopLen {
			break
		}
		resultText += string(curBit)
	}

	return resultText
}
func base64CharToSixBit(charVal uint8) [6]byte {

	for i := 0; i < len(base64Table); i++ {
		if base64Table[i] == charVal {
			resultByte := byte(i)
			var resultBytes [6]byte
			for j := 0; j < 6; j++ {
				bit := (resultByte & 0b00111111) >> 5
				resultBytes[j] = bit
				resultByte = resultByte << 1
			}
			return resultBytes
		}
	}

	return [6]byte{}
}

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func bitToBase64Encoding(bitBytes []byte) string {

	var resultTest string = ""
	var base24Padding int = 0
	var mergeBitForGetChar byte = 0
	{
		for i := 0; i < len(bitBytes); i++ {

			curBit := bitBytes[i] << (5 - i)
			mergeBitForGetChar = mergeBitForGetChar | curBit
			base24Padding++
			if i == 5 {
				resultTest += string(base64Table[mergeBitForGetChar])
				mergeBitForGetChar = 0
				i = -1
				bitBytes = bitBytes[6:]
				if base24Padding == 24 {
					base24Padding = 0
				}
			}
		}
	}

	if mergeBitForGetChar > 0 {
		mergeBitForGetChar = mergeBitForGetChar | 0b0
		resultTest += string(base64Table[mergeBitForGetChar])

		base24Padding += 6
		base24Padding = base24Padding - (base24Padding % 6)
	}

	if base24Padding > 0 {
		for paddingIdx := 0; paddingIdx < int(math.Ceil(float64((24-base24Padding)/6))); paddingIdx++ {
			resultTest += "="
		}
	}

	return resultTest
}

func binaryBitCutting(bytes []byte) []byte {

	var bitLists []byte = make([]byte, len(bytes)*8)
	idx := 0
	for idx < len(bytes) {
		val := bytes[idx]
		for i := 0; i < 8; i++ {
			bit := val >> 7
			bitLists[idx*8+i] = bit
			val = val << 1
		}
		idx++
	}

	return bitLists

}
