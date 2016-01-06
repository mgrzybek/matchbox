package api

import (
	"net/http"
)

// ignitionHandler returns a handler that responds with the ignition json the
// client machine should use.
func ignitionHandler(store Store) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		attrs := attrsFromRequest(req)
		spec, err := getMatchingSpec(store, attrs)
		if err != nil || spec.IgnitionConfig == "" {
			http.NotFound(w, req)
			return
		}

		config, err := store.IgnitionConfig(spec.IgnitionConfig)
		if err != nil {
			http.NotFound(w, req)
			return
		}
		renderJSON(w, config)
	}
	return http.HandlerFunc(fn)
}