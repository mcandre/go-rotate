package rotate

import "bytes"

const (
	UpperA = int('A')
	UpperZ = int('Z')
	LowerA = int('a')
	LowerZ = int('z')
)

func Rot13(plaintext string) string {
	return Rot(plaintext, 13)
}

func Rot(plaintext string, shift int) string {
	plainBytes := []rune(plaintext)

	var ciphertext bytes.Buffer

	for _, e := range plainBytes {
		ciphertext.WriteRune(RotRune(e, shift))
	}

	return ciphertext.String()
}

func RotRune(p rune, shift int) rune {
	b := int(p)

	switch {
	case b >= UpperA && b <= UpperZ:
		return rune((((b-UpperA)+shift)%26)+UpperA)
	case b >= LowerA && b <= LowerZ:
		return rune((((b-LowerA)+shift)%26)+LowerA)
	default:
		return p
	}
}
