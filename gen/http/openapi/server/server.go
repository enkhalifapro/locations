// Code generated by goa v3.2.5, DO NOT EDIT.
//
// openapi HTTP server
//
// Command:
// $ goa gen locations/api/design

package server

import (
	"context"
	openapi "locations/gen/openapi"
	"net/http"
	"path"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// Server lists the openapi service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the openapi service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *openapi.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"./public/swagger/", "GET", "/swagger"},
			{"./gen/http/openapi.json", "GET", "/swagger.json"},
		},
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "openapi" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
}

// Mount configures the mux to serve the openapi endpoints.
func Mount(mux goahttp.Muxer) {
	MountPublicSwagger(mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upath := path.Clean(r.URL.Path)
		rpath := upath
		if strings.HasPrefix(upath, "/swagger") {
			rpath = upath[8:]
		}
		http.ServeFile(w, r, path.Join("./public/swagger/", rpath))
	}))
	MountGenHTTPOpenapiJSON(mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/http/openapi.json")
	}))
}

// MountPublicSwagger configures the mux to serve GET request made to
// "/swagger".
func MountPublicSwagger(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/swagger/", h.ServeHTTP)
	mux.Handle("GET", "/swagger/*filepath", h.ServeHTTP)
}

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/swagger.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/swagger.json", h.ServeHTTP)
}
