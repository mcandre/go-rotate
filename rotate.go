// Package rotate provides utillity functions for Caesar shift encryption, decryption.
package rotate

import "bytes"

// ASCII code points.
const (
	UpperA = rune('A')
	UpperZ = rune('Z')
	LowerA = rune('a')
	LowerZ = rune('z')
)

// Rot13 is a symmetric encryption and decryption function for a Caesar shift of 13 on strings.
func Rot13(plaintext string) string {
	return Rot(plaintext, 13)
}

// Rot is a symmetric encryption and decryption function for an arbitrary Caesar shift on strings.
func Rot(plaintext string, shift int) string {
	plainBytes := []rune(plaintext)

	var ciphertext bytes.Buffer

	for _, e := range plainBytes {
		ciphertext.WriteRune(RotRune(e, shift))
	}

	return ciphertext.String()
}

// RotRune is a symmetric encryption and decryption function for an arbitrary Caesar shift on individual runes.
func RotRune(b rune, shift int) rune {
	switch {
	case b >= UpperA && b <= UpperZ:
		return ((b - UpperA + rune(shift)) % 26) + UpperA
	case b >= LowerA && b <= LowerZ:
		return ((b - LowerA + rune(shift)) % 26) + LowerA
	default:
		return b
	}
}
