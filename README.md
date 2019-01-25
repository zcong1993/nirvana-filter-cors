# nirvana-filter-cors [![Go Report Card](https://goreportcard.com/badge/github.com/zcong1993/nirvana-filter-cors)](https://goreportcard.com/report/github.com/zcong1993/nirvana-filter-cors)

> cors filter for nirvana framework, wrapper of [github.com/rs/cors](github.com/rs/cors)

## Usage

```go
package filters

import (
	"github.com/caicloud/nirvana/service"
	"github.com/zcong1993/nirvana-filter-cors"
)

// Filters returns a list of filters.
func Filters() []service.Filter {
	return []service.Filter{
		service.RedirectTrailingSlash(),
		service.FillLeadingSlash(),
		service.ParseRequestForm(),
		cors.AllowAll(), // add to your filter list
	}
}
```

## License

MIT &copy; zcong1993
