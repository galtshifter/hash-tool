package service

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
)

// since there is a requirement to use only standart libs, we are going to create our own stubs
type httpClientStub struct {
	stubF func(string) (*http.Response, error)
}

func (s *httpClientStub) Get(url string) (resp *http.Response, err error) {
	return s.stubF(url)
}

func ExampleService_runWorker_correct() {
	stub := httpClientStub{
		stubF: func(s string) (*http.Response, error) {
			resp := &http.Response{
				Body: ioutil.NopCloser(bytes.NewReader([]byte(s))),
			}

			return resp, nil
		},
	}
	srv := New(&stub, 5)

	input := make(chan string, 2)
	input <- "test string 1"
	input <- "test string 2"
	close(input)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	srv.runWorker(context.Background(), wg, input)

	// output:
	// test string 1 36e1d246c453fb9191bf0be36304d857
	// test string 2 3d964e44f23409ece3f028ca0643f026
}

func ExampleService_runWorker_errorHandling() {
	stub := httpClientStub{
		stubF: func(s string) (*http.Response, error) {
			return nil, errors.New("test error")
		},
	}
	srv := New(&stub, 5)

	input := make(chan string, 2)
	input <- "test string 1"
	input <- "test string 2"
	close(input)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	srv.runWorker(context.Background(), wg, input)

	// output:
	// failed to execute get request for test string 1 url: test error
	// failed to execute get request for test string 2 url: test error
}
