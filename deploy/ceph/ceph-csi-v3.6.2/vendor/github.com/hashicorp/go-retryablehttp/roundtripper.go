package retryablehttp

import (
	"net/http"
	"sync"
)

// RoundTripper implements the http.RoundTripper interface, using a retrying
// HTTP client to execute requests.
//
// It is important to note that retryablehttp doesn't always act exactly as a
// RoundTripper should. This is highly dependent on the retryable client's
// configuration.
type RoundTripper struct {
	// The client to use during requests. If nil, the default retryablehttp
	// client and settings will be used.
	Client *Client

	// once ensures that the logic to initialize the default client runs at
	// most once, in a single thread.
	once sync.Once
}

// init initializes the underlying retryable client.
func (rt *RoundTripper) init() {
	if rt.Client == nil {
		rt.Client = NewClient()
	}
}

// RoundTrip satisfies the http.RoundTripper interface.
func (rt *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	rt.once.Do(rt.init)

	// Convert the request to be retryable.
	retryableReq, err := FromRequest(req)
	if err != nil {
		return nil, err
	}

	// Execute the request.
	return rt.Client.Do(retryableReq)
}
