package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/soypita/go-workshop/internals/api"
	"github.com/soypita/go-workshop/internals/api/mocks"
	"github.com/soypita/go-workshop/internals/handler"
)

func TestSimpleHandler_Hello(t *testing.T) {
	tests := []struct {
		name     string
		joke     string
		err      error
		codeWant int
		bodyWant string
	}{
		{
			name:     "simple test",
			joke:     "test joke",
			err:      nil,
			codeWant: 200,
			bodyWant: "test joke",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiMock := &mocks.Client{}
			apiMock.On("GetJoke").Return(&api.JokeResponse{
				Joke: tt.joke,
			}, tt.err)

			h := handler.NewSimpleHandler(apiMock)

			req, _ := http.NewRequest("GET", "/hello", nil)
			rr := httptest.NewRecorder()

			h.Hello(rr, req)

			// handler := http.HandlerFunc(h.Hello)
			// handler.ServeHTTP(rr, req)

			gotRaw, _ := ioutil.ReadAll(rr.Body)
			got := string(gotRaw)
			if got != tt.bodyWant {
				t.Errorf("Response body wrong %s want %s", got, tt.bodyWant)
			}

			if status := rr.Result().StatusCode; status != tt.codeWant {
				t.Errorf("Response status wrong %d want %d", status, tt.codeWant)
			}
		})
	}
}
