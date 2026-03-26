package tiktokads

import (
	"errors"
	"fmt"
	"strings"
)

type errNotFound struct {
	Type string
}

var searchTimeoutError = errors.New("search-timeout")
var campaignNotFoundError = errors.New("campaign-not-found")
var adNotFoundError = errors.New("ad-not-found")

func newErrNotFound(objectType string) error {
	return &errNotFound{
		Type: objectType,
	}
}

func (e *errNotFound) Error() string {
	return fmt.Sprintf("%s-not-found", strings.ToLower(e.Type))
}
