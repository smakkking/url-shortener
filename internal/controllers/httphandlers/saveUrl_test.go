package httphandlers

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/smakkking/url-shortener/internal/services"
	mock_services "github.com/smakkking/url-shortener/internal/services/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHandler_SaveURL(t *testing.T) {
	type mockBehaviour func(s *mock_services.MockStorage)

	testTable := []struct {
		name                string
		inputBody           string
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"url": "https://www.youtube.com/dsafsdf"}`,
			mockBehaviour: func(s *mock_services.MockStorage) {
				inputURL, _ := url.Parse("https://www.youtube.com/dsafsdf")
				s.EXPECT().SaveURL(context.Background(), "1234567890", *inputURL)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"status": "OK", "url": "http://localhost:8080/1234567890"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock_services.NewMockStorage(ctrl)
			testCase.mockBehaviour(m)

			service := services.NewService(m)
			handler := NewHandler(service)

			// test server
			srv := http.NewServeMux()
			srv.HandleFunc("/create", handler.SaveURL)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(
				http.MethodPost,
				"/create",
				bytes.NewBufferString(testCase.inputBody),
			)

			srv.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}
