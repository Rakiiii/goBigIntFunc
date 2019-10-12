package bigintlib

import "math/big"
import "testing"

func TestPow(t *testing.T) {
	b := big.NewInt(int64(2))
	c := Pow(b, int64(10))

	if c.Cmp(big.NewInt(int64(1024))) != 0 {
		t.Error("Expected 1024, got", c.String())
	}
}

func TestInc(t *testing.T) {
	b := big.NewInt(100)
	Inc(b)
	if b.Cmp(big.NewInt(101)) != 0 {
		t.Error("Expected 101 , got ", b.String())
	}

}
