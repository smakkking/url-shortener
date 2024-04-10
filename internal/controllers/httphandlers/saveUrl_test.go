package httphandlers

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/smakkking/url-shortener/internal/services"
	mock_services "github.com/smakkking/url-shortener/internal/services/mocks"
	"github.com/smakkking/url-shortener/pkg/keygenerator"
	"github.com/stretchr/testify/assert"
)

func TestHandler_SaveURL(t *testing.T) {
	type mockBehaviour func(ctx context.Context, s *mock_services.MockStorage)

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
			mockBehaviour: func(ctx context.Context, s *mock_services.MockStorage) {
				inputURL, _ := url.Parse("https://www.youtube.com/dsafsdf")
				s.EXPECT().SaveURL(ctx, "1234567890", *inputURL).Return("1234567890", nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"status":"OK","url":"http://localhost:8080/1234567890"}` + "\n",
		},
		{
			name:                "Invalid json",
			inputBody:           `345kj3k5l`,
			mockBehaviour:       func(ctx context.Context, s *mock_services.MockStorage) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"status":"Error","error":"invalid json"}` + "\n",
		},
		{
			name:                "Incorrect url",
			inputBody:           `{"url": "     %%2"}`,
			mockBehaviour:       func(ctx context.Context, s *mock_services.MockStorage) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"status":"Error","error":"incorrect url"}` + "\n",
		},
		{
			name:      "Cant save url",
			inputBody: `{"url": "https://www.youtube.com/dsafsdf"}`,
			mockBehaviour: func(ctx context.Context, s *mock_services.MockStorage) {
				inputURL, _ := url.Parse("https://www.youtube.com/dsafsdf")
				s.EXPECT().SaveURL(ctx, "1234567890", *inputURL).Return("", errors.New("can't save url"))
			},
			expectedStatusCode:  http.StatusInternalServerError,
			expectedRequestBody: `{"status":"Error","error":"can't save url"}` + "\n",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock_services.NewMockStorage(ctrl)

			service := services.NewService(m, keygenerator.FixedKeyGenerator{Key: "1234567890"})
			handler := NewHandler(service)

			// test server
			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/create",
				bytes.NewBufferString(testCase.inputBody),
			)

			rctx := chi.NewRouteContext()
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
			testCase.mockBehaviour(r.Context(), m)

			handler.SaveURL(w, r)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}
