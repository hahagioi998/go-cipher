package cipher_test

import (
	"github.com/ituknown/go-cipher"
	"testing"
)

func TestGenerateBcryptSecret(t *testing.T) {

	var plainPwd = "1q2w3e4r!@#"

	var bcryptSecret = cipher.GenerateBcryptSecret(plainPwd, 0)

	t.Logf("Password Plain: %s\n", plainPwd)
	t.Logf("BcryptSecret: %s\n", bcryptSecret)
}

func TestVerifyBcryptSecret(t *testing.T) {

	var plainPwd = "1q2w3e4r!@#"

	var bcryptSecret = cipher.GenerateBcryptSecret(plainPwd, 0)

	t.Logf("Password Plain: %s\n", plainPwd)
	t.Logf("BcryptSecret: %s\n", bcryptSecret)

	bol := cipher.VerifyBcryptSecret(bcryptSecret, plainPwd)

	t.Logf("Verify: %t\n", bol)
}
