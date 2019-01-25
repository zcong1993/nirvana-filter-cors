package cors

import (
	"github.com/caicloud/nirvana/service"
	cors2 "github.com/rs/cors"
	"net/http"
)

type corsWrapper struct {
	*cors2.Cors
	optionPassthrough bool
}

func (c corsWrapper) build() service.Filter {
	return func(w http.ResponseWriter, req *http.Request) bool {
		c.HandlerFunc(w, req)
		// for preflight request, end with 204
		if !c.optionPassthrough &&
			req.Method == http.MethodOptions &&
			req.Header.Get("Access-Control-Request-Method") != "" {
			w.WriteHeader(http.StatusNoContent)
			return false
		}
		// if optionPassthrough is true, handle response by user, so should pass through
		return true
	}
}

// New is filter wrapper with provided options
func New(options cors2.Options) service.Filter {
	return corsWrapper{Cors: cors2.New(options), optionPassthrough: options.OptionsPassthrough}.build()
}

// AllowAll is filter wrapper of cors.AllowAll()
func AllowAll() service.Filter {
	return corsWrapper{Cors: cors2.AllowAll()}.build()
}

// Default is filter wrapper of cors.Default()
func Default() service.Filter {
	return corsWrapper{Cors: cors2.Default()}.build()
}
