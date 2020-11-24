package utils

import "strings"

func ConstructQuery(queryValues map[string]string, fields []string) map[string]string {
	if queryValues == nil {
		queryValues = make(map[string]string)
	}

	if fields != nil && len(fields) > 0 {
		queryValues["fields"] = strings.Join(fields, ",")
	}

	return queryValues
}
