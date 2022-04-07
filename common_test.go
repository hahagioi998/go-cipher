package cipher_test

import (
	"github.com/ituknown/go-cipher"
	"testing"
)

func TestScanPwd(t *testing.T) {

	var plainPwd = cipher.ScanPwd()
	t.Logf("scan pwd: %s\n", plainPwd)
}
