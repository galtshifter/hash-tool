package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/galtshifter/hash-tool/internal/service"
)

const (
	httpClientTimeout = 5 * time.Second
)

var parallelReqLimit int

func init() {
	flag.IntVar(&parallelReqLimit, "parallel", 10, "number of max parallel requests")
	flag.Parse()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS interrupts
	go func() {
		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

		s := <-signalCh
		signal.Stop(signalCh)

		fmt.Printf("receive OS signal, cancel work, signal: %v\n", s)
		cancel()
	}()

	args := flag.Args()
	reqURLs, err := prepareURLs(args)
	if err != nil {
		fmt.Printf("failed to prepare request urls: %v\n", err)
		return
	}

	httpClient := &http.Client{
		Timeout: httpClientTimeout,
	}

	srv := service.New(httpClient, parallelReqLimit)

	srv.Run(ctx, reqURLs)
}
