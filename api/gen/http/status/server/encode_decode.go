// Code generated by goa v3.2.0, DO NOT EDIT.
//
// status HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	"context"
	"net/http"

	status "github.com/tektoncd/hub/api/gen/status"
	goahttp "goa.design/goa/v3/http"
)

// EncodeStatusResponse returns an encoder for responses returned by the status
// Status endpoint.
func EncodeStatusResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*status.StatusResult)
		enc := encoder(ctx, w)
		body := NewStatusResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}