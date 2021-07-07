package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"sync"
)

func (s *Service) runWorker(ctx context.Context, wg *sync.WaitGroup, input <-chan string) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case reqURL, ok := <-input:
			if !ok {
				return
			}

			resp, err := s.httpClient.Get(reqURL)
			if err != nil {
				fmt.Printf("failed to execute get request for %s url: %v\n", reqURL, err)
				continue
			}
			defer resp.Body.Close()

			respData, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("failed to read response body: %v\n", err)
				continue
			}

			respHash := md5.Sum(respData)

			fmt.Printf("%s %x\n", reqURL, respHash)
		}
	}
}
