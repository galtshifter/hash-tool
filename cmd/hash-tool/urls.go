package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func prepareURLs(args []string) ([]string, error) {
	res := make([]string, 0, len(args))
	for _, v := range args {
		v = strings.TrimPrefix(v, "http://")
		v = strings.TrimPrefix(v, "https://")
		v = "http://" + v

		u, err := url.ParseRequestURI(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %s: %v", v, err)
		}

		if u.Host == "" {
			return nil, errors.New(v + ": empty host")
		}

		if !strings.Contains(u.Host, ".") {
			return nil, errors.New(v + ": no top level domain")
		}

		res = append(res, u.String())
	}

	return res, nil
}
