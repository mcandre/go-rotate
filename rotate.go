// Package rotate provides utillity functions for Caesar shift encryption, decryption.
package rotate

import "bytes"

// Version is semver.
var Version = "0.0.2"

// ASCII code points.
const (
	UpperA = int('A')
	UpperZ = int('Z')
	LowerA = int('a')
	LowerZ = int('z')
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
func RotRune(p rune, shift int) rune {
	b := int(p)

	switch {
	case b >= UpperA && b <= UpperZ:
		return rune((((b - UpperA) + shift) % 26) + UpperA)
	case b >= LowerA && b <= LowerZ:
		return rune((((b - LowerA) + shift) % 26) + LowerA)
	default:
		return p
	}
}
