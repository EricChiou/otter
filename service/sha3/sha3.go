package sha3

import (
	"otter/config"
	"otter/pkg/sha3"
)

// Encrypt sha3 encrypt
func Encrypt(pwd string, lenght ...int) string {
	if len(lenght) > 0 {
		return sha3.Encrypt(pwd, lenght[0])
	}
	return sha3.Encrypt(pwd, config.Sha3Len)
}
