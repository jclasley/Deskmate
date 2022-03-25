// Package dds contains the methods for interacting with the DDS.
//
// Because Deskmate does not contain any PII, this package exports a
// single method -- ListenToDDS() -- that should be run in a goroutine
// that polls the DDS and then simply submits back a response without
// doing any actual work.
package dds

import "net/http"

const ServiceUUID = "4a6c7dd8-71d9-4655-a8d0-8d0fb6e6c7b2"

// init function registers the service to the
// DDS, even if the service is already registered.
func init() {

}

func formatRequest(r *http.Request) {
	r.Header.Set("User-Agent", "circleci/itd-deskmate")
	r.Header.Set("X-Service-Id", ServiceUUID)

}