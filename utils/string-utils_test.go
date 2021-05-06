package utils

import (
	"fmt"
	"testing"
)

func TestSecureRandomAlphaString(t *testing.T) {
	fmt.Println(SecureRandomAlphaString(16))
}
