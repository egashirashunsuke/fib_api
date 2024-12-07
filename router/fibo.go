package router

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FiboHandler(c echo.Context) error {
	nStr := c.QueryParam("n")
	if nStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "400",
			"message": "Query parameter 'n' is required.",
		})
	}

	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "400",
			"message": "Query parameter 'n' must be a positive integer.",
		})
	}

	result, err := Fibocal(n)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "400",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": result.String(),
	})
}

func Fibocal(n int) (*big.Int, error) {
	if n <= 0 {
		return nil, fmt.Errorf("n must be a positive integer")
	}
	if n == 1 || n == 2 {
		return big.NewInt(1), nil
	}

	//uint64ではn=99が処理できなかったため、big型に変更
	a := big.NewInt(1)
	b := big.NewInt(1)
	for i := 3; i <= n; i++ {
		c := new(big.Int).Add(a, b)
		a.Set(b)
		b.Set(c)
	}
	return b, nil
}
