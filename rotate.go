// Package rotate provides utillity functions for Caesar shift encryption, decryption.
package rotate

import "bytes"

// Version is semver.
var Version = "0.0.2"

// ASCII code points.
const (
	upperA = rune('A')
	upperZ = rune('Z')
	lowerA = rune('a')
	lowerZ = rune('z')
)

// Rot13 is a symmetric encryption and decryption function for a Caesar shift of 13 on strings.
func Rot13(plaintext string) string {
	return Rot(plaintext, 13)
}

// Rot is a symmetric encryption and decryption function for an arbitrary Caesar shift on strings.
func Rot(plaintext string, shift int) string {
	var ciphertext bytes.Buffer

	for _, e := range plaintext {
		ciphertext.WriteRune(RotRune(e, shift))
	}

	return ciphertext.String()
}

// RotRune is a symmetric encryption and decryption function for an arbitrary Caesar shift on individual runes.
func RotRune(b rune, shift int) rune {
	switch {
	case b >= upperA && b <= upperZ:
		return ((b - upperA + rune(shift)) % 26) + upperA
	case b >= lowerA && b <= lowerZ:
		return ((b - lowerA + rune(shift)) % 26) + lowerA
	default:
		return b
	}
}
