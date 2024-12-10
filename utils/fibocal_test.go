package utils

import (
	"testing"
)

func TestFibocal(t *testing.T) {
	tests := []struct {
		n        int
		expected string // big.Intを文字列で比較(big.Intは構造体であるため)
		hasError bool
	}{
		{n: 1, expected: "1", hasError: false},
		{n: 2, expected: "1", hasError: false},
		{n: 10, expected: "55", hasError: false},
		{n: 99, expected: "218922995834555169026", hasError: false},
		{n: 0, expected: "", hasError: true},
		{n: -5, expected: "", hasError: true},
	}

	calculator := RealFiboCalculator{}

	//t.Runを使うことで，テストを階層化
	for _, test := range tests {
		t.Run("Fibocal", func(t *testing.T) {
			result, err := calculator.Fibocal(test.n)

			if test.hasError {
				if err == nil {
					t.Errorf("n=%dのとき、エラーが期待されるがnilだった", test.n)
				}
			} else {
				if err != nil {
					t.Errorf("n=%dのとき、エラーが期待されないがエラーが起きた: %v", test.n, err)
				}

				if result != test.expected {
					t.Errorf("n=%dのとき、 %sが期待されるが %sがでた", test.n, test.expected, result)
				}
			}
		})
	}
}
