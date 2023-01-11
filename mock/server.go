package mock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/julienschmidt/httprouter"
	"github.com/neticdk/go-bitbucket/bitbucket"
)

type MockBackendOption func(*httprouter.Router)

func NewMockServer(opts ...MockBackendOption) *httptest.Server {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		WriteError(w, http.StatusNotFound, []bitbucket.ErrorMessage{{Message: "Not Found"}})
	})
	for _, o := range opts {
		o(router)
	}
	mockServer := httptest.NewServer(router)
	return mockServer
}

type errorResponse struct {
	Errors []bitbucket.ErrorMessage `json:"errors"`
}

func WriteError(w http.ResponseWriter, status int, errs []bitbucket.ErrorMessage) {
	w.WriteHeader(status)
	err := &errorResponse{Errors: errs}
	json.NewEncoder(w).Encode(err)
}

type EndpointPattern struct {
	Pattern string
	Method  string
}

func WithRequestMatchHandler(ep EndpointPattern, handler http.Handler) MockBackendOption {
	return func(router *httprouter.Router) {
		router.Handler(ep.Method, ep.Pattern, handler)
	}
}

func WithRequestMatch(ep EndpointPattern, responsesFIFO ...interface{}) MockBackendOption {
	responses := [][]byte{}

	for _, r := range responsesFIFO {
		switch v := r.(type) {
		case []byte:
			responses = append(responses, v)
		default:
			b, err := json.Marshal(r)
			if err != nil {
				panic(fmt.Sprintf("go-bitbucket/mock: unable to serialiaze json: %v", err))
			}
			responses = append(responses, b)
		}
	}

	return WithRequestMatchHandler(ep, &FIFOReponseHandler{
		Responses: responses,
	})
}

type FIFOReponseHandler struct {
	lock         sync.Mutex
	Responses    [][]byte
	CurrentIndex int
}

func (h *FIFOReponseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.lock.Lock()
	defer h.lock.Unlock()

	if h.CurrentIndex > len(h.Responses) {
		panic(fmt.Sprintf("go-bitbucket/mock: no more mocks available for %s", r.URL.Path))
	}

	defer func() {
		h.CurrentIndex++
	}()

	w.Write(h.Responses[h.CurrentIndex])
}
