package octo

import (
	"time"

	"github.com/willabides/octo-go/options"
)

//go:generate go run ./generator -schema "api.github.com.json" -pkgpath "github.com/willabides/octo-go" -pkg octo

// RelLink is the name for a relative link in a Link header. Used for paging.
type RelName string

// Common RelLink values
const (
	RelNext  RelName = "next"
	RelPrev  RelName = "prev"
	RelFirst RelName = "first"
	RelLast  RelName = "last"
)

// ISOTimeString returns a pointer to tm formated as an iso8601/rfc3339 string
func ISOTimeString(tm time.Time) *string {
	return String(tm.Format(time.RFC3339))
}

// String returns a pointer to s
func String(s string) *string {
	return &s
}

// Bool returns a pointer to b
func Bool(b bool) *bool {
	return &b
}

// Int64 returns a pointer to i
func Int64(i int64) *int64 {
	return &i
}

// Client is a set of options to apply to requests
type Client []options.Option

// NewClient returns a new Client
func NewClient(opt ...options.Option) Client {
	return opt
}
