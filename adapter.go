package web

import "net/http"

// Adapter wraps a http.Handler and returns a http.Handler. For a detailed
// discussion please checkout:
// https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
type Adapter func(http.Handler) http.Handler

// Wrap wraps the actual http.Handler that handles the request with a list of
// adapters. The adapters provided are called in sequence and if all adapters
// forward the call to the next one the handler will be called.
func Wrap(handler http.Handler, adapters ...Adapter) http.Handler {
	for i := len(adapters) - 1; i >= 0; i-- {
		// Wrap the handler with the adapter and use result to chain to the
		// next adapter.
		handler = adapters[i](handler)
	}

	// Return the handler returned by the first adapter in the list. This will
	// ensure that the handlers are called in the order of invocation.
	return handler
}
