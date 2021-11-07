package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"
)

//PrettyJSON from object
func PrettyJSON(obj interface{}) string {
	jsonBytes, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// JSON from object
func JSON(obj interface{}) string {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// MergeStringToStringMap, last one wins on conflick
func MergeStringToStringMap(ms ...map[string]string) map[string]string {
	res := map[string]string{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

func ToMapStringToString(m map[string]interface{}) map[string]string {
	var found = make(map[string]string)

	for key, value := range m {
		if value == nil {
			continue
		}
		switch t := value.(type) {
		case *uint:
			found[key] = strconv.FormatUint(uint64(*t), 10)
		case *uint8:
			found[key] = strconv.FormatUint(uint64(*t), 10)
		case *int:
			found[key] = strconv.FormatInt(int64(*t), 10)
		case *int16:
			found[key] = strconv.FormatInt(int64(*t), 10)
		case *int32:
			found[key] = strconv.FormatInt(int64(*t), 10)
		case *int64:
			found[key] = strconv.FormatInt(int64(*t), 10)
		case *float32:
			found[key] = strconv.FormatFloat(float64(*t), 'f', 32, 32)
		case *float64:
			found[key] = strconv.FormatFloat(float64(*t), 'f', 64, 64)
		case *bool:
			found[key] = strconv.FormatBool(*t)
		case *string:
			found[key] = *t
		}

	}
	return found
}
func ResolveEndpointUrl(basePath string, path string) (*url.URL, error) {
	parts := strings.Split(path, "?")
	if len(parts) > 2 {
		// only allow one ? in the path
		return nil, errors.New("invalid path")
	}
	pathSegment := joinURL(basePath, parts[0])
	if len(parts) == 2 {
		pathSegment = fmt.Sprintf("%s?%s", pathSegment, parts[1])
	}
	t, _ := url.Parse(pathSegment)
	t = t.ResolveReference(t)
	return t, nil
}

//joinURL joins a base with url paths
func joinURL(base string, paths ...string) string {
	p := path.Join(paths...)
	return fmt.Sprintf("%s/%s", strings.TrimRight(base, "/"), strings.TrimLeft(p, "/"))
}
