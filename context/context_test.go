package context

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type subStore struct {
	response string
}

func (s subStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	data := "Hello, World!"
	svr := Server(&subStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}
