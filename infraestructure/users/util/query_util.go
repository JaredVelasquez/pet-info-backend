package util

import (
	"fmt"
	"regexp"
)

func ReplaceNamedParams(query string, params map[string]interface{}) (string, []interface{}) {
	var args []interface{}
	var re = regexp.MustCompile(`:\w+`)

	usedParams := make(map[string]int)
	i := 1

	query = re.ReplaceAllStringFunc(query, func(param string) string {
		if _, exists := usedParams[param]; !exists {
			usedParams[param] = i
			i++
		}
		args = append(args, params[param[1:]])
		return fmt.Sprintf("$%d", usedParams[param])
	})

	return query, args
}
