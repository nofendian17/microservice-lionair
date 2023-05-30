package parser

import "net/http"

type Parser func(response *http.Response) (string, error)
