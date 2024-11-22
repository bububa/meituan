package model

import (
	"net/url"
)

// GetRequest get request interface
type GetRequest interface {
	// Encode encode request to string
	Values(values url.Values)
}
