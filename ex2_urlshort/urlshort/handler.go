package urlshort

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// MapHandler configures the given fallback handler for the given url paths as map data structure
func MapHandler(paths map[string]string, fallback *mux.Router) *mux.Router {
	for route, url := range paths {
		fallback.HandleFunc(route, func(res http.ResponseWriter, req *http.Request) {
			http.Redirect(res, req, url, http.StatusMovedPermanently)
		})
	}
	return fallback
}

// JSONHandler configures the given fallback handler for the given url paths as JSON data structure
func JSONHandler(jsonPaths []byte, fallback *mux.Router) (*mux.Router, error) {
	var urlMap map[string]string

	if err := json.Unmarshal(jsonPaths, &urlMap); err != nil {
		return nil, err
	}

	return MapHandler(urlMap, fallback), nil
}
