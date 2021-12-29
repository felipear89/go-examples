package books

import (
	"errors"
	"github.com/felipear89/go-examples/pkg/api/router"
	"github.com/felipear89/go-examples/pkg/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx/fxtest"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/felipear89/go-examples/pkg/model"
	repo "github.com/felipear89/go-examples/pkg/repository/books"
)

func Test_getBooks(t *testing.T) {
	type args struct {
		getAll repo.GetAll
		route  string
	}
	tests := []struct {
		name         string
		args         args
		expectedCode int
		expectedBody string
	}{
		{
			name: "get all books",
			args: args{
				getAll: func() ([]*model.Book, error) { return []*model.Book{}, nil },
				route:  "/api/books",
			},
			expectedCode: 200,
			expectedBody: "{\"items\":[]}",
		},
		{
			name: "should return error",
			args: args{
				getAll: func() ([]*model.Book, error) { return nil, errors.New("test error") },
				route:  "/api/books",
			},
			expectedCode: 500,
			expectedBody: "{\"error\":\"test error\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := config.NewServer(fxtest.NewLifecycle(t), &config.Config{})
			routes := router.NewRouteAPI(app)
			Register(routes, Params{
				GetAll: tt.args.getAll,
			})
			req, _ := http.NewRequest("GET", tt.args.route, nil)
			res, err := app.Test(req, -1)

			assert.Equalf(t, tt.expectedCode, res.StatusCode, tt.name)

			body, err := ioutil.ReadAll(res.Body)

			assert.Nilf(t, err, tt.name)
			assert.Equalf(t, tt.expectedBody, string(body), tt.name)
		})
	}
}
