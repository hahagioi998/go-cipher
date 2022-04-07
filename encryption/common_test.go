package encryption_test

import (
	"testing"

	"ituknown.com/cipher/encryption"
)

func TestScanPwd(t *testing.T) {
	var plainPwd string = encryption.ScanPwd()
	t.Logf("scan pwd: %s\n", plainPwd)
}
