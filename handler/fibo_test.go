package handler_test

import (
	"fibo_api/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/assert"
)

//モック化に挑戦しようとしたが間に合わなかったので、
//外部関数であるutils.Fibocal(n)を直接使用してテスト

func TestFiboHandler(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		expected   string
		statusCode int
	}{
		{"ValidInput", "?n=10", `{"result":"55"}`, http.StatusOK},
		{"MissingParam", "", `{"status":"400","message":"クエリパラメータが必要です"}`, http.StatusBadRequest},
		{"InvalidParam", "?n=abc", `{"status":"400","message":"クエリパラメータは0以上です"}`, http.StatusBadRequest},
		{"NegativeParam", "?n=-5", `{"status":"400","message":"クエリパラメータは0以上です"}`, http.StatusBadRequest},
		{"MinInput", "?n=1", `{"result":"1"}`, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/fibo"+tt.query, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, handler.FiboHandler(c)) {
				assert.Equal(t, tt.statusCode, rec.Code)
				assert.JSONEq(t, tt.expected, rec.Body.String())
			}
		})
	}
}
