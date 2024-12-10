package utils

import (
	"fmt"
	"math/big"
)

type RealFiboCalculator struct{}

// n番目のフィボナッチ数列を計算する関数
func (r *RealFiboCalculator) Fibocal(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("nが0以下です")
	}
	if n == 1 || n == 2 {
		return "1", nil
	}

	//uint64ではn=99が処理できなかったため、big型に変更
	a := big.NewInt(1)
	b := big.NewInt(1)
	for i := 3; i <= n; i++ {
		c := new(big.Int).Add(a, b)
		a.Set(b)
		b.Set(c)
	}
	return b.String(), nil
}
