package bigintlib

import "math/big"

func Pow(num *big.Int, pow int64) *big.Int {
	res := big.NewInt(1)
	for i := int64(0); i < pow; i++ {
		res.Mul(res, num)
	}
	return res
}

func Inc(num *big.Int) {
	num.Add(num, big.NewInt(1))
}
