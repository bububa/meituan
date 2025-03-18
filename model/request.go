package model

import (
	"io"
	"net/url"
)

// GetRequest get request interface
type GetRequest interface {
	// Encode encode request to string
	Values(values url.Values)
}

type PostRequest interface {
	Encode(w io.Writer)
}
