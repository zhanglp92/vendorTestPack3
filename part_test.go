package vendorTestPack3

import "testing"

func TestPack3(t *testing.T) {
	t.Logf(NewVendorTestPack3(1, 2).String())
}
