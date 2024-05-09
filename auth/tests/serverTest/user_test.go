package server

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/Prokopevs/ccc/auth/internal/core"
	"github.com/Prokopevs/ccc/auth/internal/server"

	mock_server "github.com/Prokopevs/ccc/auth/internal/server/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

func TestServer_me(t *testing.T) {
	type mockBehavior func(r *mock_server.MockService, ctx context.Context, initData string, inviterId int)

	tests := []struct {
		name                 string
		ctx                  context.Context
		headerName           string
		headerValue          string
		inviterId            int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			ctx:         context.Background(),
			headerName:  "initData",
			headerValue: "hejjdsje",
			inviterId:   0,
			mockBehavior: func(r *mock_server.MockService, ctx context.Context, headerValue string, inviterId int) {
				r.EXPECT().GetUserInfo(ctx, headerValue, inviterId).Return(&core.UserInfo{Id: 1, Firstname: "sr", Username: "yf"}, core.CodeOK, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"firstname":"sr","username":"yf"}`,
		},
		{
			name:        "Invalid Header Name",
			ctx:         context.Background(),
			headerName:  "initDa",
			headerValue: "hejjdsje",
			inviterId:   0,
			mockBehavior: func(r *mock_server.MockService, ctx context.Context, headerValue string, inviterId int) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"code":"NO_HEADER","errorInfo":"no initData"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)

			repo := mock_server.NewMockService(c)
			test.mockBehavior(repo, test.ctx, test.headerValue, test.inviterId)

			server := &server.HTTP{service: repo}

			// Init Endpoint
			r := gin.New()
			r.GET("/me", server.me)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/me", nil)
			req.Header.Set(test.headerName, test.headerValue)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)

			defer c.Finish()
		})
	}
}
