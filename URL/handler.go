package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		dest, ok := pathsToUrls[path]
		if ok { //if dest has been found, then ok
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}
	mappath := buildMap(pathUrls) //make map from paths and directly use on maphandler to re-use code
	return MapHandler(mappath, fallback), nil
}

func buildMap(pathUrls []pathUrl) map[string]string {
	mappath := make(map[string]string)
	for _, pu := range pathUrls {
		mappath[pu.Path] = pu.URL //make map from paths
	}
	return mappath
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil //won't get an error here for sure, so can pass nil
}

//added handler for JSON files which works similarly to that of YAMLhandler

func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseJSON(jsonBytes)
	if err != nil {
		return nil, err
	}
	mappath := make(map[string]string)
	for _, pu := range pathUrls {
		mappath[pu.Path] = pu.URL
	}
	return MapHandler(mappath, fallback), nil //make map from paths and directly use on maphandler to re-use code
}

func parseJSON(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := json.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}

	return pathUrls, nil //won't get error, pass nil
}
