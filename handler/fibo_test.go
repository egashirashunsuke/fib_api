package handler_test

import (
	"fibo_api/handler"
	"fibo_api/mocks"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestFiboHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//t.Runを使うことで，テストを階層化
	tests := []struct {
		name         string
		mockSetup    func(mockCalculator *mocks.MockFiboCalculator)
		expectedCode int
		expectedBody string
	}{
		{
			name: "正常系: 正しい値を返す",
			mockSetup: func(mockCalculator *mocks.MockFiboCalculator) {
				mockCalculator.EXPECT().Fibocal(10).Return("55", nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"result":"55"}`,
		},
		{
			name: "異常系: エラーを返す",
			mockSetup: func(mockCalculator *mocks.MockFiboCalculator) {
				mockCalculator.EXPECT().Fibocal(10).Return("", fmt.Errorf("計算エラー"))
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"status":"400","message":"計算エラー"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックインスタンスを生成
			mockCalculator := mocks.NewMockFiboCalculator(ctrl)

			// モックの動作をセットアップ
			tt.mockSetup(mockCalculator)

			// ハンドラにモックを注入
			h := &handler.Handler{Calculator: mockCalculator}

			// テスト用のHTTPリクエストを作成
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/fibo?n=10", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// ハンドラを実行
			if assert.NoError(t, h.FiboHandler(c)) {
				// ステータスコードの確認
				assert.Equal(t, tt.expectedCode, rec.Code)
				// レスポンスの確認
				assert.JSONEq(t, tt.expectedBody, rec.Body.String())
			}
		})
	}
}
