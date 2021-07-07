package service

import (
	"context"
	"net/http"
	"sync"
)

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type Service struct {
	httpClient  httpClient
	workerLimit int
}

func New(httpClient httpClient, workerLimit int) *Service {
	return &Service{
		httpClient:  httpClient,
		workerLimit: workerLimit,
	}
}

func (s *Service) Run(ctx context.Context, reqEndpoints []string) {
	input := make(chan string, len(reqEndpoints))
	for _, v := range reqEndpoints {
		input <- v
	}
	close(input)

	wg := new(sync.WaitGroup)

	wg.Add(s.workerLimit)
	for i := 0; i < s.workerLimit; i++ {
		go s.runWorker(ctx, wg, input)
	}

	wg.Wait()
}
