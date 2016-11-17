package rotate_test

import (
	"testing"
	"testing/quick"

	"github.com/mcandre/go-rotate"
)

func TestSymmetricROT13Encryption(t *testing.T) {
	symmetricProperty := func(s string) bool {
		return rotate.Rot13(rotate.Rot13(s)) == s
	}

	err := quick.Check(symmetricProperty, nil)

	if err != nil {
		t.Error(err)
	}
}
