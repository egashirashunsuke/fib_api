package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// FiboCalculatorインターフェース
type FiboCalculator interface {
	Fibocal(n int) (string, error)
}

// Handler構造体
type Handler struct {
	Calculator FiboCalculator
}

func (h *Handler) FiboHandler(c echo.Context) error {
	nStr := c.QueryParam("n")
	if nStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "400",
			"message": "クエリパラメータが必要です",
		})
	}

	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "400",
			"message": "クエリパラメータは0以上です",
		})
	}

	result, err := h.Calculator.Fibocal(n)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "400",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": result,
	})
}
