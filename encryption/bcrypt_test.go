package encryption_test

import (
	"testing"

	"ituknown.com/cipher/encryption"
)

func TestGenerateBcryptSecret(t *testing.T) {

	var plainPwd string = "1q2w3e4r!@#"

	var bcryptSecret string = encryption.GenerateBcryptSecret(plainPwd, 0)

	t.Logf("Password Plain: %s\n", plainPwd)
	t.Logf("BcryptSecret: %s\n", bcryptSecret)
}
