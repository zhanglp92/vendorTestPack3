package vendorTestPack3

import (
	"github.com/zhanglp92/vendorTestPack1"
	"github.com/zhanglp92/vendorTestPack2"
)

// VendorTestPack3 the type for test vendor,
// create type3.
type VendorTestPack3 struct {
	v1 *vendorTestPack1.VendorTestPack1
	v2 *vendorTestPack2.VendorTestPack2
}

// NewVendorTestPack3 is the function for create VendorTestPack3.
func NewVendorTestPack3(n1, n2 int) *VendorTestPack3 {
	v := VendorTestPack3{
		v1: vendorTestPack1.NewVendorTestPack1(n1),
		v2: vendorTestPack2.NewVendorTestPack2(n2),
	}
	return &v
}

func (v VendorTestPack3) String() string {
	return "1, 2"
	// return fmt.Sprintf("%v", v)
}
